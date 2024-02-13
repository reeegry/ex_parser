import json
from sdamgia import SdamGIA


subject = "math"
sdamgia = SdamGIA()

with open("categores.json", "w") as file:
    json.dump(sdamgia.get_catalog(subject), file)
