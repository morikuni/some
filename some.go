package some

type Some struct{}

func (s Some) Int() IntSpec         { return Int }
func (s Some) Int64() Int64Spec     { return Int64 }
func (s Some) Uint() UintSpec       { return Uint }
func (s Some) Float64() Float64Spec { return Float64 }
func (s Some) Float32() Float32Spec { return Float32 }
func (s Some) Bool() BoolSpec       { return Bool }
func (s Some) String() StringSpec   { return String }
func (s Some) Time() TimeSpec       { return Time }
func (s Some) URL() URLSpec         { return URL }
