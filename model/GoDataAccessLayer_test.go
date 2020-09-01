package model

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type Person struct {
	Name           string
	Address        string
	Age            int
	Male           bool
	Married        bool
	Interests      []string
	GraduationDate time.Time
	Friends        []*Person
	Spouse         *Person
	Children       map[string]*Person
	Pet            Pet
}

func (p *Person) IncreaseAge() {
	p.Age++
}

func (p *Person) IsOld() bool {
	return p.Age > 40
}

type Pet interface {
	Name() string
	GetKind() string
	GetAge() int
	Cubs() []Pet
}

type Cat struct {
	name string
	age  int
	cubs []*Cat
}

func (c *Cat) Name() string {
	return c.name
}
func (c *Cat) GetKind() string {
	return "Cat"
}
func (c *Cat) GetAge() int {
	return c.age
}
func (c *Cat) Cubs() []Pet {
	ret := make([]Pet, len(c.cubs))
	for _, cub := range c.cubs {
		ret = append(ret, cub)
	}
	return ret
}

func MakeTestPerson() *Person {
	pet := &Cat{
		name: "Luca",
		age:  3,
		cubs: []*Cat{
			&Cat{
				name: "Yuri",
				age:  1,
				cubs: nil,
			},
		},
	}
	return &Person{
		Name:           "James",
		Address:        "21 Jump Street",
		Age:            25,
		Married:        true,
		Male:           true,
		Interests:      []string{"Football", "Game", "Coding"},
		GraduationDate: time.Date(2005, time.July, 23, 12, 0, 0, 0, time.UTC),
		Friends: []*Person{
			&Person{
				Name:           "Johnson",
				Address:        "Pinewood Road",
				Age:            23,
				Male:           true,
				Interests:      []string{"Swimming", "Hiking", "Party"},
				GraduationDate: time.Date(2006, time.July, 23, 12, 0, 0, 0, time.UTC),
			},
			&Person{
				Name:           "Peter",
				Address:        "Metro Complex",
				Age:            21,
				Male:           true,
				GraduationDate: time.Date(2007, time.July, 23, 12, 0, 0, 0, time.UTC),
			},
		},
		Spouse: &Person{
			Name:           "Lynda",
			Address:        "21 Jump Street",
			Age:            23,
			Married:        true,
			Male:           false,
			GraduationDate: time.Date(2008, time.July, 23, 12, 0, 0, 0, time.UTC),
			Friends: []*Person{
				&Person{
					Name:           "Lucy",
					Address:        "Hilbury Blvrd",
					Age:            23,
					Male:           false,
					GraduationDate: time.Date(2009, time.July, 23, 12, 0, 0, 0, time.UTC),
				},
				&Person{
					Name:           "Darla",
					Address:        "Low Road Passing Blvrd",
					Age:            21,
					Male:           false,
					GraduationDate: time.Date(2010, time.July, 23, 12, 0, 0, 0, time.UTC),
				},
			},
			Spouse:   nil,
			Children: nil,
			Pet:      nil,
		},
		Children: map[string]*Person{
			"Christen": &Person{
				Name: "Christen",
				Age:  3,
				Male: false,
			},
			"Graham": &Person{
				Name: "Graham",
				Age:  1,
				Male: true,
			},
		},
		Pet: pet,
	}
}

func TestGoValueNode_Array(t *testing.T) {
	person := MakeTestPerson()
	actorNode := NewGoValueNode(reflect.ValueOf(person), "actor")

	friendsNode, err := actorNode.GetChildNodeByField("Friends")
	assert.NoError(t, err)
	assert.True(t, friendsNode.IsArray())
	typ, err := friendsNode.GetArrayType()
	assert.NoError(t, err)
	assert.Equal(t, "*model.Person", typ.String())

	interestsNode, err := actorNode.GetChildNodeByField("Interests")
	assert.NoError(t, err)
	assert.True(t, interestsNode.IsArray())
	typ, err = interestsNode.GetArrayType()
	assert.NoError(t, err)
	assert.Equal(t, reflect.String, typ.Kind())

	arrLen, err := friendsNode.Length()
	assert.NoError(t, err)
	assert.Equal(t, 2, arrLen)
	val0, err := friendsNode.GetArrayValueAt(0)
	assert.NoError(t, err)
	assert.Equal(t, reflect.Ptr, val0.Kind())
	assert.Equal(t, "model.Person", val0.Elem().Type().String())
	friends0Node, err := friendsNode.GetChildNodeByIndex(0)
	assert.NoError(t, err)
	assert.Equal(t, "actor.Friends[0]", friends0Node.IdentifiedAs())
	friends0NameNode, err := friends0Node.GetChildNodeByField("Name")
	assert.NoError(t, err)
	assert.Equal(t, "actor.Friends[0].Name", friends0NameNode.IdentifiedAs())
	assert.True(t, friends0NameNode.IsString())
	val, err := friends0NameNode.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "Johnson", val.String())
}

func TestGoValueNode_Map(t *testing.T) {
	person := MakeTestPerson()
	actorNode := NewGoValueNode(reflect.ValueOf(person), "actor")
	childrenNode, err := actorNode.GetChildNodeByField("Children")
	assert.NoError(t, err)
	assert.True(t, childrenNode.IsMap())
	childrenChristenValue, err := childrenNode.GetMapValueAt(reflect.ValueOf("Christen"))
	assert.NoError(t, err)
	assert.Equal(t, "*model.Person", childrenChristenValue.Type().String())
	childrenChristenNode, err := childrenNode.GetChildNodeBySelector(reflect.ValueOf("Christen"))
	assert.NoError(t, err)
	assert.True(t, childrenChristenNode.IsObject())
	t.Logf(childrenChristenNode.IdentifiedAs())
}

func TestGoValueNode_Array_Set(t *testing.T) {
	person := MakeTestPerson()
	actorNode := NewGoValueNode(reflect.ValueOf(person), "actor")
	interestNode, err := actorNode.GetChildNodeByField("Interests")
	assert.NoError(t, err)
	assert.True(t, interestNode.IsArray())

	interest1Node, err := interestNode.GetChildNodeByIndex(1)
	assert.NoError(t, err)
	assert.True(t, interest1Node.IsString())
	val, err := interest1Node.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "Game", val.String())

	err = interestNode.SetArrayValueAt(1, reflect.ValueOf("Gaming"))
	assert.NoError(t, err)

	interest1Node, err = interestNode.GetChildNodeByIndex(1)
	assert.NoError(t, err)
	assert.True(t, interest1Node.IsString())
	val, err = interest1Node.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "Gaming", val.String())

	l, err := interestNode.Length()
	assert.NoError(t, err)
	assert.Equal(t, 3, l)

	err = interestNode.AppendValue([]reflect.Value{reflect.ValueOf("Diving")})
	assert.NoError(t, err)

	l, err = interestNode.Length()
	assert.NoError(t, err)
	assert.Equal(t, 4, l)

	interest1Node, err = interestNode.GetChildNodeByIndex(3)
	assert.NoError(t, err)
	assert.True(t, interest1Node.IsString())
	val, err = interest1Node.GetValue()
	assert.NoError(t, err)
	assert.Equal(t, "Diving", val.String())

}

func TestGoValueNodeSetValues(t *testing.T) {
	person := MakeTestPerson()
	actorNode := NewGoValueNode(reflect.ValueOf(person), "actor")
	valueNode, err := actorNode.GetChildNodeByField("Address")
	assert.NoError(t, err)
	assert.True(t, valueNode.IsString())
	addressValue, err := valueNode.GetValue()
	assert.Equal(t, "21 Jump Street", addressValue.String())

	err = actorNode.SetObjectValueByField("Address", reflect.ValueOf("22 Dunk Street"))
	assert.NoError(t, err)
	valueNode, err = actorNode.GetChildNodeByField("Address")
	assert.NoError(t, err)
	assert.True(t, valueNode.IsString())
	addressValue, err = valueNode.GetValue()
	assert.Equal(t, "22 Dunk Street", addressValue.String())

	// set value of different type
	err = actorNode.SetObjectValueByField("Address", reflect.ValueOf(22))
	assert.Error(t, err)

	// set value of non existent field
	err = actorNode.SetObjectValueByField("NonExistent", reflect.ValueOf("22 Dunk Street"))
	assert.Error(t, err)
}

func TestGoValueNode(t *testing.T) {
	person := MakeTestPerson()
	actorNode := NewGoValueNode(reflect.ValueOf(person), "actor")

	// check initial identifiedAs
	assert.Equal(t, "actor", actorNode.IdentifiedAs())
	assert.False(t, actorNode.HasParent())

	// check on age field
	actorAgeNode, err := actorNode.GetChildNodeByField("Age")
	assert.NoError(t, err, "got %s", err)
	assert.True(t, actorAgeNode.HasParent())
	assert.Equal(t, "actor.Age", actorAgeNode.IdentifiedAs())

	intValue, err := actorAgeNode.GetValue()
	assert.NoError(t, err)
	assert.True(t, intValue.Kind() == reflect.Int)

	assert.True(t, actorAgeNode.IsInteger())
	assert.False(t, actorAgeNode.IsObject())
	assert.False(t, actorAgeNode.IsString())
	assert.False(t, actorAgeNode.IsArray())
	assert.False(t, actorAgeNode.IsMap())
	assert.False(t, actorAgeNode.IsBool())
	assert.False(t, actorAgeNode.IsReal())
	assert.False(t, actorAgeNode.IsTime())

	// check on name
	actorNameNode, err := actorNode.GetChildNodeByField("Name")
	assert.NoError(t, err, "got %s", err)
	assert.True(t, actorNameNode.HasParent())
	assert.Equal(t, "actor.Name", actorNameNode.IdentifiedAs())
	stringValue, err := actorNameNode.GetValue()
	assert.NoError(t, err)
	assert.True(t, stringValue.Kind() == reflect.String)
	assert.False(t, actorNameNode.IsInteger())
	assert.False(t, actorNameNode.IsObject())
	assert.True(t, actorNameNode.IsString())
	assert.False(t, actorNameNode.IsArray())
	assert.False(t, actorNameNode.IsMap())
	assert.False(t, actorNameNode.IsBool())
	assert.False(t, actorNameNode.IsReal())
	assert.False(t, actorNameNode.IsTime())

	// check on name
	actorMarriedNode, err := actorNode.GetChildNodeByField("Married")
	assert.NoError(t, err, "got %s", err)
	assert.True(t, actorMarriedNode.HasParent())
	assert.Equal(t, "actor.Married", actorMarriedNode.IdentifiedAs())
	boolValue, err := actorMarriedNode.GetValue()
	assert.NoError(t, err)
	assert.True(t, boolValue.Kind() == reflect.Bool)
	assert.False(t, actorMarriedNode.IsInteger())
	assert.False(t, actorMarriedNode.IsObject())
	assert.False(t, actorMarriedNode.IsString())
	assert.False(t, actorMarriedNode.IsArray())
	assert.False(t, actorMarriedNode.IsMap())
	assert.True(t, actorMarriedNode.IsBool())
	assert.False(t, actorMarriedNode.IsReal())
	assert.False(t, actorMarriedNode.IsTime())

	actorGraduationNode, err := actorNode.GetChildNodeByField("GraduationDate")
	assert.NoError(t, err, "got %s", err)
	assert.True(t, actorGraduationNode.HasParent())
	assert.Equal(t, "actor.GraduationDate", actorGraduationNode.IdentifiedAs())
	assert.False(t, actorGraduationNode.IsInteger())
	assert.True(t, actorGraduationNode.IsObject())
	assert.False(t, actorGraduationNode.IsString())
	assert.False(t, actorGraduationNode.IsArray())
	assert.False(t, actorGraduationNode.IsMap())
	assert.False(t, actorGraduationNode.IsBool())
	assert.False(t, actorGraduationNode.IsReal())
	assert.True(t, actorGraduationNode.IsTime())

}

func TestConstantFunctionCalls(t *testing.T) {
	textNode := NewGoValueNode(reflect.ValueOf("   SomeWithSpace  "), "SpacedText")
	retVal, err := textNode.CallFunction("Trim")
	assert.NoError(t, err)
	assert.Equal(t, "string", retVal.Type().String())
	assert.Equal(t, "SomeWithSpace", retVal.String())
}
