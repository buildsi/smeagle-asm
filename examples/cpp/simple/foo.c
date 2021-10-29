#include <cstdio>
#include <ostream>
#include <iostream>
#include <string>
#include "assert.h"
#include "foo.h"

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }}) {

{{ .Function | AssertArgs }}
}
