import json
from sdamgia import SdamGIA


subjects = ["inf", "hist", "math", "mathb", "chem", "rus", "bio", "en", "geo", "de", "soc", "fr", "lit", "sp"]
sdamgia = SdamGIA()

for subject in subjects:
    with open(f"categores_{subject}.json", "w") as file:
        print(subject)
        try:
            json.dump(sdamgia.get_catalog(subject), file)
        except:
            print(f"can't create {subject} json")
