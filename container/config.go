package container

import (
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"

	containerreflect "github.com/cosmos/cosmos-sdk/container/reflect"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type config struct {
	autoGroupTypes   map[reflect.Type]bool
	onePerScopeTypes map[reflect.Type]bool

	// logging
	loggers   []func(string)
	indentStr string

	// graphing
	graphviz      *graphviz.Graphviz
	graph         *cgraph.Graph
	visualizers   []func(string)
	logVisualizer bool
}

func newConfig() (*config, error) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return nil, err
	}

	return &config{
		autoGroupTypes:   map[reflect.Type]bool{},
		onePerScopeTypes: map[reflect.Type]bool{},
		graphviz:         g,
		graph:            graph,
	}, nil
}

func (c *config) indentLogger() {
	c.indentStr = c.indentStr + " "
}

func (c *config) dedentLogger() {
	if len(c.indentStr) > 0 {
		c.indentStr = c.indentStr[1:]
	}
}

func (c config) logf(format string, args ...interface{}) {
	s := fmt.Sprintf(c.indentStr+format, args...)
	for _, logger := range c.loggers {
		logger(s)
	}
}

func (c *config) generateGraph() {
	buf := &bytes.Buffer{}
	err := c.graphviz.Render(c.graph, graphviz.XDOT, buf)
	if err != nil {
		c.logf("Error rendering DOT graph: %+v", err)
		return
	}

	dot := buf.String()
	if c.logVisualizer {
		c.logf("DOT Graph: %s", dot)
	}

	for _, v := range c.visualizers {
		v(dot)
	}

	err = c.graph.Close()
	if err != nil {
		c.logf("Error closing graph: %+v", err)
		return
	}

	err = c.graphviz.Close()
	if err != nil {
		c.logf("Error closing graphviz: %+v", err)
	}
}

func (c *config) addFuncVisualizer(f func(string)) {
	c.visualizers = append(c.visualizers, func(dot string) {
		f(dot)
	})
}

func (c *config) enableLogVisualizer() {
	c.logVisualizer = true
}

func (c *config) addFileVisualizer(filename string, format string) {
	c.visualizers = append(c.visualizers, func(_ string) {
		err := c.graphviz.RenderFilename(c.graph, graphviz.Format(format), filename)
		if err != nil {
			c.logf("Error saving graphviz file %s with format %s: %+v", filename, format, err)
		} else {
			path, err := filepath.Abs(filename)
			if err == nil {
				c.logf("Saved graph of container to %s", path)
			}
		}
	})
}

func (c *config) locationGraphNode(location containerreflect.Location) (*cgraph.Node, error) {
	node, found, err := c.findOrCreateGraphNode(location.Name())
	if err != nil {
		return nil, err
	}

	if found {
		return node, nil
	}

	node = node.SetShape(cgraph.BoxShape)
	node.SetColor("lightgrey")
	return node, nil
}

func (c *config) typeGraphNode(typ reflect.Type) (*cgraph.Node, error) {
	node, found, err := c.findOrCreateGraphNode(typ.String())
	if err != nil {
		return nil, err
	}

	if found {
		return node, nil
	}

	node.SetColor("lightgrey")
	return node, err
}

func (c *config) findOrCreateGraphNode(name string) (node *cgraph.Node, found bool, err error) {
	node, err = c.graph.Node(name)
	if err != nil {
		return nil, false, err
	}

	if node != nil {
		return node, true, nil
	}

	node, err = c.graph.CreateNode(name)
	return node, false, err
}

func (c *config) addGraphEdge(from *cgraph.Node, to *cgraph.Node) {
	_, err := c.graph.CreateEdge("", from, to)
	if err != nil {
		c.logf("error creating graph edge")
	}
}
