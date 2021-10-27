#pragma once

#include <cstdint>
#include <string>

// Structs used in the function should be declared first
{{ .Function | DeclareStructs }} 

void {{ .Function | GetFunctionName }}({{ .Function | AsFormalParams }});
