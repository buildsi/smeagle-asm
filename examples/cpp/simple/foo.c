#include <cstdio>
#include <ostream>
#include <iostream>
#include <string>
#include "foo.h"

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }}) {

{{ .Function | PrintArgs }}
}
