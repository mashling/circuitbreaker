package circuitbreaker

import (
	"math/rand"
	"testing"
	"time"
	"math"
)

func TestCircuitBreakerModeA(t *testing.T) {
	rand.Seed(1)
	clock := time.Unix(1533930608, 0)
	now = func() time.Time {
		now := clock
		clock = clock.Add(time.Duration(rand.Intn(2)+1) * time.Second)
		return now
	}
	defer func() {
		now = time.Now
	}()


	execute := func(serviceName string,values map[string]interface{}, should error) {
		factory := Factory{}
		breaker, err := factory.Make(serviceName, values)
		if err != nil {
			t.Fatal(err)
		}
		if values == nil {
			values = make(map[string]interface{})
		}
		values["context"] = "testA"
		err = breaker.UpdateRequest(values)
		if err != nil {
			t.Fatal(err)
		}

		err = breaker.Execute()
		if err != should {
			t.Fatalf("error should be %v but is %v", should, err)
		}
	}

	for i := 0; i < 4; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	execute("reset",nil, nil)
	execute("reset",map[string]interface{}{"operation": "reset"}, nil)

	for i := 0; i < 5; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
	execute("reset",map[string]interface{}{"operation": "counter"}, nil)

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
}


func TestCircuitBreakerModeB(t *testing.T) {
	rand.Seed(1)
	clock := time.Unix(1533930608, 0)
	now = func() time.Time {
		now := clock
		clock = clock.Add(time.Duration(rand.Intn(2)+1) * time.Second)
		return now
	}
	defer func() {
		now = time.Now
	}()

	execute := func(serviceName string,values map[string]interface{}, should error) {
		factory := Factory{}
		breaker, err := factory.Make(serviceName, values)
		if err != nil {
			t.Fatal(err)
		}
		if values == nil {
			values = make(map[string]interface{})
		}
		values["context"] = "testB"
		values["mode"] = CircuitBreakerModeB

		err = breaker.UpdateRequest(values)
		if err != nil {
			t.Fatal(err)
		}
		err = breaker.Execute()
		if err != should {
			t.Fatalf("error should be %v but is %v", should, err)
		}
	}

	for i := 0; i < 4; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	clock = clock.Add(60 * time.Second)

	for i := 0; i < 5; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
	execute("reset",map[string]interface{}{"operation": "counter"}, nil)

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
}

func TestCircuitBreakerModeC(t *testing.T) {
	rand.Seed(1)
	clock := time.Unix(1533930608, 0)
	now = func() time.Time {
		now := clock
		clock = clock.Add(time.Duration(rand.Intn(2)+1) * time.Second)
		return now
	}
	defer func() {
		now = time.Now
	}()

	execute := func(serviceName string,values map[string]interface{}, should error) {
		factory := Factory{}
		breaker, err := factory.Make(serviceName, values)
		if err != nil {
			t.Fatal(err)
		}
		if values == nil {
			values = make(map[string]interface{})
		}
		values["context"] = "testC"
		values["mode"] = CircuitBreakerModeC

		err = breaker.UpdateRequest(values)
		if err != nil {
			t.Fatal(err)
		}
		err = breaker.Execute()
		if err != should {
			t.Fatalf("error should be %v but is %v", should, err)
		}
	}

	for i := 0; i < 4; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	clock = clock.Add(60 * time.Second)

	for i := 0; i < 4; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	execute("reset",nil, nil)
	execute("reset",map[string]interface{}{"operation": "reset"}, nil)

	for i := 0; i < 5; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "counter"}, nil)
	}

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
	execute("reset",map[string]interface{}{"operation": "counter"}, nil)

	execute("reset",nil, ErrorCircuitBreakerTripped)

	clock = clock.Add(60 * time.Second)

	execute("reset",nil, nil)
}

func TestCircuitBreakerModeD(t *testing.T) {
	rand.Seed(1)
	clock := time.Unix(1533930608, 0)
	now = func() time.Time {
		now := clock
		clock = clock.Add(time.Duration(rand.Intn(2)+1) * time.Second)
		return now
	}
	defer func() {
		now = time.Now
	}()

	execute := func(serviceName string, values map[string]interface{}, should error) error {
		factory := Factory{}
		breaker, err := factory.Make(serviceName, values)
		if err != nil {
			t.Fatal(err)
		}
		if values == nil {
			values = make(map[string]interface{})
		}
		values["context"] = "testD"
		values["mode"] = CircuitBreakerModeD

		err = breaker.UpdateRequest(values)
		if err != nil {
			t.Fatal(err)
		}
		err = breaker.Execute()
		if err != should {
			t.Fatalf("error should be %v but is %v", should, err)
		}
		return err
	}

	for i := 0; i < 100; i++ {
		execute("reset",nil, nil)
		execute("reset",map[string]interface{}{"operation": "reset"}, nil)
	}
	p := circuitBreakerContexts.GetContext("testD", 5).Probability(now())
	if math.Floor(p*100) != 0.0 {
		t.Fatalf("probability should be zero but is %v", math.Floor(p*100))
	}

	type Test struct {
		a, b error
	}
	tests := []Test{
		{nil, nil},
		{nil, nil},
		{ErrorCircuitBreakerTripped, nil},
		{ErrorCircuitBreakerTripped, nil},
		{nil, nil},
		{ErrorCircuitBreakerTripped, nil},
		{ErrorCircuitBreakerTripped, nil},
		{ErrorCircuitBreakerTripped, nil},
	}
	for _, test := range tests {
		err := execute("reset",nil, test.a)
		if err != nil {
			continue
		}
		execute("reset",map[string]interface{}{"operation": "counter"}, test.b)
	}

	tests = []Test{
		{nil, nil},
		{nil, nil},
		{nil, nil},
		{nil, nil},
		{nil, nil},
	}
	for _, test := range tests {
		err := execute("reset",nil, test.a)
		if err != nil {
			continue
		}
		execute("reset",map[string]interface{}{"operation": "reset"}, test.b)
	}
	p = circuitBreakerContexts.GetContext("testD", 5).Probability(now())
	if math.Floor(p*100) != 0.0 {
		t.Fatalf("probability should be zero but is %v", math.Floor(p*100))
	}
}