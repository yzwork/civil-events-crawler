// Package gen_test contains tests for the gen package
package gen_test

import (
	"bytes"
	"github.com/joincivil/civil-events-crawler/pkg/gen"
	"strings"
	"testing"
	"time"
)

var (
	event1 = &gen.EventHandler{
		EventMethod: "Application",
		EventName:   "_Application",
		EventType:   "CivilTCRContractApplication",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
			{Type: "common.Address"},
		},
	}
	event2 = &gen.EventHandler{
		EventMethod: "ApplicationRemoved",
		EventName:   "_ApplicationRemoved",
		EventType:   "CivilTCRContractApplicationRemoved",
		ParamValues: []*gen.EventHandlerMethodParam{
			{Type: "common.Address"},
		},
	}
	testWatchers = &gen.ContractData{
		PackageName:         "watcher",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandler{
			event1,
			event2,
		},
	}
	testFilterers = &gen.ContractData{
		PackageName:         "retrieve",
		ContractImportPath:  "github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandler{
			event1,
			event2,
		},
	}
	testBadGoFmtWatchers = &gen.ContractData{
		PackageName:         "watcher",
		ContractImportPath:  "((((github.com/joincivil/civil-events-crawler/pkg/generated/contract",
		ContractTypePackage: "contract",
		ContractTypeName:    "CivilTCRContract((((((",
		GenTime:             time.Now().UTC(),
		EventHandlers: []*gen.EventHandler{
			event1,
			event2,
		},
	}
)

func TestNameToContractType(t *testing.T) {
	names := gen.NameToContractTypes.Names()
	if len(names) != len(gen.NameToContractTypes) {
		t.Errorf("Len of Names() should be same as length of NameToContractTypes")
	}
	namesFound := 0
	for _, name := range names {
		if _, ok := gen.NameToContractTypes[name]; ok {
			namesFound++
		}
	}
	if namesFound != len(gen.NameToContractTypes) {
		t.Errorf("Matching names found should be same num as in NameToContractTypes")
	}
}

func TestGenerateEventHandlersFromBadTemplates(t *testing.T) {
	bufWatcher := &bytes.Buffer{}
	err := gen.GenerateEventHandlersFromTemplate(bufWatcher, testWatchers, true,
		"badhandlerName")
	if err == nil {
		t.Errorf("Should have failed with a bad handler name: err: %v", err)
	}
	err = gen.GenerateEventHandlersFromTemplate(bufWatcher, testBadGoFmtWatchers, true,
		"watcher")
	if err == nil {
		t.Errorf("Should have failed with a gofmt failure: err: %v", err)
	}
}

func TestGenerateEventHandlersFromTemplate(t *testing.T) {
	bufWatcher := &bytes.Buffer{}
	bufFilterer := &bytes.Buffer{}
	err := gen.GenerateEventHandlersFromTemplate(bufWatcher, testWatchers, true, "watcher")
	if err != nil {
		t.Errorf("Error generating watchers: err: %v", err)
	}
	err = gen.GenerateEventHandlersFromTemplate(bufFilterer, testFilterers, true, "filterer")
	if err != nil {
		t.Errorf("Error generating filterers: err: %v", err)
	}

	// TODO(PN or IS): Some basic checks, need more here.
	watcherCode := bufWatcher.String()
	filtererCode := bufFilterer.String()

	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) StartWatchers") {
		t.Error("Did not see expected Startwatchers in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated watcher code")
	}
	if !strings.Contains(watcherCode, "func (w *CivilTCRContractWatchers) startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated watcher code")
	}
	if !strings.Contains(filtererCode, "func (r *CivilTCRContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated Filterer code")
	}
	if !strings.Contains(filtererCode, "func (r *CivilTCRContractFilterers) startFilterApplication") {
		t.Error("Did not see expected startFilterApplication in the generated filterer code")
	}

}

const BadContractType gen.ContractType = 500

func TestGenerateWatchersForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, BadContractType, "watcher", "watcher")
	if err == nil {
		t.Errorf("Should have failed when given an unknown contract type: err: %v", err)
	}

	err = gen.GenerateCivilEventHandlers(buf, gen.CivilTcrContractType, "watcher", "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Civil TCR contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartWatchers") {
		t.Error("Did not see expected StartWatchers in the generated watcher code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) StartCivilTCRContractWatchers") {
		t.Error("Did not see expected StartCivilTCRContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) startWatchApplication") {
		t.Error("Did not see expected startWatchApplication in the generated code")
	}
	if !strings.Contains(code, "func (w *CivilTCRContractWatchers) startWatchApplicationRemoved") {
		t.Error("Did not see expected startWatchApplicationRemoved in the generated code")
	}
}

func TestGenerateWatchersForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.NewsroomContractType, "watcher", "watcher")
	if err != nil {
		t.Errorf("Error generating watchers for the Newsroom contract: err: %v", err)
	}

	// TODO(PN): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartWatchers") {
		t.Error("Did not see expected StartWatchers in the generated watcher code")
	}
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) StartNewsroomContractWatchers") {
		t.Error("Did not see expected StartNewsroomContractWatchers in the generated code")
	}
	if !strings.Contains(code, "func (w *NewsroomContractWatchers) startWatchRevisionUpdated") {
		t.Error("Did not see expected startWatchRevisionUpdated in the generated code")
	}
}

func TestGenerateRetrieversForCivilTcr(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.CivilTcrContractType, "filterer", "filterer")
	if err != nil {
		t.Errorf("Error generating filterers for the Civil TCR contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (r *CivilTCRContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated code")
	}
	if !strings.Contains(code, "func (r *CivilTCRContractFilterers) startFilterApplication") {
		t.Error("Did not see expected startFilterApplication in the generated code")
	}
	if !strings.Contains(code, "func (r *CivilTCRContractFilterers) startFilterApplicationRemoved") {
		t.Error("Did not see expected startFilterApplicationRemoved in the generated code")
	}
}

func TestGenerateRetrieversForNewsroom(t *testing.T) {
	buf := &bytes.Buffer{}
	err := gen.GenerateCivilEventHandlers(buf, gen.NewsroomContractType, "filterer", "filterer")
	if err != nil {
		t.Errorf("Error generating retrievers for the Newsroom contract: err: %v", err)
	}

	// TODO(IS): Do some sanity check tests.  ADD MORE!
	code := buf.String()
	if !strings.Contains(code, "func (r *NewsroomContractFilterers) StartFilterers") {
		t.Error("Did not see expected StartFilterers in the generated code")
	}
	if !strings.Contains(code, "func (r *NewsroomContractFilterers) startFilterRevisionUpdated") {
		t.Error("Did not see expected startFilterRevisionUpdated in the generated code")
	}
}
