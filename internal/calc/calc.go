package calc

// Event send data to broker
type Event struct {
	Event string    `json:"event"`
	Args  EventArgs `json:"args"`
}

// EventArgs contains received data.
type EventArgs struct {
	A int `json:"a"`
	B int `json:"b"`
}

// EventResult contains result data.
type EventResult struct {
	Result float32 `json:"result"`
}

// Add function for sum.
func Add(a int, b int) float32 {
	return float32(a + b)
}

// Sub function for subtraction.
func Sub(a int, b int) float32 {
	return float32(a - b)
}

// Mut function for multiplication.
func Mut(a int, b int) float32 {
	return float32(a * b)
}

// Divide two numbers.
func Divide(a int, b int) float32 {
	return float32(a) / float32(b)
}
