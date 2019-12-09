// Package launchpad file runtime.go contains runtime support for reference
// tracking and output assembled object.
// for evaluation hierarchy and finalized assembled object tracking.
package launchpad

import (
	"fmt"
	"log"
	"strings"
)

// assembledOrg enables runtime reference tracking and allows resources to organized into an organization.
type assembledOrg struct {
	// resourceRegistry is a map of refId to resource for all processed resources.
	resourceRegistry map[string]resourceHandler
	// referenceTracker is a map of refId to referrer list for all processed resources.
	referenceTracker map[string][]resourceHandler
	// org represents the assembled organization based on processed resources.
	org orgYAML
}

// String implements Stringer and generates a string representation.
func (ao *assembledOrg) String() string { return strings.Join(ao.dump(0), "\n") }

// assembleResourcesToOrg takes in resources and assembles into an organization.
func assembleResourcesToOrg(rs []resourceHandler) *assembledOrg {
	ao := assembledOrg{}
	ao.resourceRegistry = make(map[string]resourceHandler)
	ao.referenceTracker = make(map[string][]resourceHandler)

	// discover resources in a DFS style
	// updating ao.resourceRegistry and ao.referenceTracker for reference resolution
	for _, r := range rs {
		if err := r.addToOrg(&ao); err != nil {
			log.Fatalln("error validating YAML", err.Error())
		}
	}
	// assemble each discovered resource onto a finalized org
	if err := ao.resolveReferences(); err != nil {
		log.Fatalln("unable to resolve referenceTracker between YAML resources", err.Error())
	}
	return &ao
}

// registerResource registers a resource onto resourceRegistry and referenceTracker for later resolution.
//
// registerResource will add new entry in assembledOrg.resourceRegistry for refId to resource. If
// the resource being registered (src) has a reference (dst) to another resource, it will
// be registered onto the assembledOrg.referenceTracker for dst -> [src...] mapping.
//
// If there are conflicting resources of the same id, silently pick the latest.
func (ao *assembledOrg) registerResource(src resourceHandler, dst *referenceYAML) {
	srcId := src.refId()
	ao.resourceRegistry[srcId] = src // silently replace mapping if exist

	if dst == nil {	// no reference
		return
	}

	if dst.TargetType() == Organization { // Initialize Organization
		if err := ao.org.initializeByRef(dst); err != nil { // attempt to initialize org
			log.Fatalln(fmt.Sprintf("org already initialized as %s, cannot initialize to %s",
				ao.org.Spec.Id, dst.TargetId))
		}
		// Org is special that we need to manually register it's ID for them
		ao.resourceRegistry[ao.org.refId()] = &ao.org
		// TODO take care of explicit org initialization case
	}

	// update referenceTracker for references from src
	ao.referenceTracker[dst.refId()] = append(ao.referenceTracker[dst.refId()], src)
}

// resolveReferences loops through referenceTracker to link up references.
//
// referenceTracker relies on resourceRegistry to determine if referenced object exist.
func (ao *assembledOrg) resolveReferences() error {
	for k, v := range ao.referenceTracker {
		target, found := ao.resourceRegistry[k]
		if !found {
			log.Fatalln(fmt.Sprintf("reference %s does not exist", k))
		}
		// each resource holds its own resolve logic
		if err := target.resolveReferences(v); err != nil {
			return err
		}
	}
	return nil
}

// dump generates debug string slices representation.
func (ao *assembledOrg) dump(ind int) []string {
	indent := strings.Repeat(" ", ind)

	buff := []string{fmt.Sprintf("%sReference Maps:", indent)}
	for k, v := range ao.referenceTracker {
		var srcs []string
		for _, vv := range v {
			srcs = append(srcs, vv.refId())
		}
		buff = append(buff, fmt.Sprintf("%s  * %s <- [%s]", indent, k, strings.Join(srcs, ", ")))
	}

	// Retrieve resources infos
	buff = append(buff, fmt.Sprintf("%sReferencable Targets:", indent))
	for k := range ao.resourceRegistry {
		buff = append(buff, fmt.Sprintf("%s  - %s", indent, k))
	}

	// Retrieve org info
	buff = append(buff, ao.org.dump(ind)...)
	return buff
}
