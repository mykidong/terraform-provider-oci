// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v43/common"
)

// MacroField The type representing the macro field concept. Macro fields have an expression to define a macro.
type MacroField struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	Expr *Expression `mandatory:"false" json:"expr"`

	Type BaseType `mandatory:"false" json:"type"`

	// Specifies whether the type of macro fields is inferred from an expression or useType (false) or the source field (true).
	IsUseSourceType *bool `mandatory:"false" json:"isUseSourceType"`

	UseType *ConfiguredType `mandatory:"false" json:"useType"`

	// Labels are keywords or labels that you can add to data assets, dataflows, and so on. You can define your own labels and use them to categorize content.
	Labels []string `mandatory:"false" json:"labels"`
}

//GetKey returns Key
func (m MacroField) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m MacroField) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m MacroField) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m MacroField) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m MacroField) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m MacroField) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m MacroField) GetDescription() *string {
	return m.Description
}

func (m MacroField) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m MacroField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeMacroField MacroField
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeMacroField
	}{
		"MACRO_FIELD",
		(MarshalTypeMacroField)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *MacroField) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key             *string          `json:"key"`
		ModelVersion    *string          `json:"modelVersion"`
		ParentRef       *ParentReference `json:"parentRef"`
		ConfigValues    *ConfigValues    `json:"configValues"`
		ObjectStatus    *int             `json:"objectStatus"`
		Name            *string          `json:"name"`
		Description     *string          `json:"description"`
		Expr            *Expression      `json:"expr"`
		Type            basetype         `json:"type"`
		IsUseSourceType *bool            `json:"isUseSourceType"`
		UseType         *ConfiguredType  `json:"useType"`
		Labels          []string         `json:"labels"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Name = model.Name

	m.Description = model.Description

	m.Expr = model.Expr

	nn, e = model.Type.UnmarshalPolymorphicJSON(model.Type.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Type = nn.(BaseType)
	} else {
		m.Type = nil
	}

	m.IsUseSourceType = model.IsUseSourceType

	m.UseType = model.UseType

	m.Labels = make([]string, len(model.Labels))
	for i, n := range model.Labels {
		m.Labels[i] = n
	}

	return
}