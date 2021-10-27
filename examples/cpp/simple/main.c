#include "foo.h"
#include <iostream>
#include <string>

int main() {

     // Initialize each formal param
{{ .Function | DeclareArgs }}

     // bigcall(1, 2, 3, 4, 5, bigthing);
     {{ .Function | CallFunction }}
}
