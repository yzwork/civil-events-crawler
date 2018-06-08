// Package gen contains all the components for code generation.
package gen

const filtererTmpl = `
// Code generated by 'gen/watchergen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'listener/watchergen' for more details
// File was generated at {{.GenTime}}
package {{.PackageName}}

import (
    log "github.com/golang/glog"
    "fmt"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"

    "github.com/joincivil/civil-events-crawler/pkg/model"
{{if .ContractImportPath -}}
    "{{.ContractImportPath}}"
{{- end}}
{{if .AdditionalImports -}}
{{- range .AdditionalImports}}
    "{{.}}"
{{- end}}
{{- end}}
)

type {{.ContractTypeName}}Filterers struct {}

func (r *{{.ContractTypeName}}Filterers) ContractName() string {
    return "{{.ContractTypeName}}"
}

func (r *{{.ContractTypeName}}Filterers) StartFilterers(client bind.ContractBackend, contractAddress common.Address,
    pastEvents *[]model.CivilEvent, startBlock uint64) error {
    return r.Start{{.ContractTypeName}}Filterers(client, contractAddress, pastEvents, startBlock)
}

// Start{{.ContractTypeName}}Filterers retrieves events for {{.ContractTypeName}}
func (r *{{.ContractTypeName}}Filterers) Start{{.ContractTypeName}}Filterers(client bind.ContractBackend, 
    contractAddress common.Address, pastEvents *[]model.CivilEvent, startBlock uint64) error {
    contract, err := {{.ContractTypePackage}}.New{{.ContractTypeName}}(contractAddress, client)
    if err != nil {
        log.Errorf("Error initializing Start{{.ContractTypeName}}: err: %v", err)
        return err
    }

    var opts = &bind.FilterOpts{
        Start: startBlock,
    }

{{if .EventHandlers -}}
{{- range .EventHandlers}}

    err = startFilter{{.EventMethod}}(opts, contract, pastEvents)
    if err != nil {
        return fmt.Errorf("Error retrieving {{.EventMethod}}: err: %v", err)
    }

{{- end}}
{{- end}}

    return nil
}

{{if .EventHandlers -}}
{{- range .EventHandlers}}

func startFilter{{.EventMethod}}(opts *bind.FilterOpts, _contract *{{$.ContractTypePackage}}.{{$.ContractTypeName}}, pastEvents *[]model.CivilEvent) error {
    itr, err := _contract.Filter{{.EventMethod}}(
        opts,
    {{- if .ParamValues -}}
    {{range .ParamValues}}
        []{{.Type}}{},
    {{- end}}
    {{end}}
    )
    if err != nil {
        log.Errorf("Error getting event {{.EventMethod}}: %v", err)
        return err
    }
    nextEvent := itr.Next()
    for nextEvent {
        civilEvent := model.NewCivilEvent("{{.EventMethod}}", itr.Event)
        *pastEvents = append(*pastEvents, *civilEvent)
        nextEvent = itr.Next()
    }
    return nil
}

{{- end}}
{{- end}}
`
