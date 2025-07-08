# gt-items
tool for reading and searching growtopia `items.dat` with advanced filtering.

---

![image](https://github.com/user-attachments/assets/a10c040a-b735-421a-a96f-42856b7cca88)

---

##  item searching examples

| command | description |
|------------------------------------------------------------|----------------------------------------------------------|
| `search --id=100:200`                                      | get items with IDs between 100 and 200                  |
| `search --rarity=!999`                                     | get items whose rarity is **not** 999                   |
| `search --name=dirt`                                       | get items whose name contains "dirt"                    |
| `search --type=storage\|passwordstorage`                   | get items with type `Storage` **or** `PasswordStorage`  |
| `search --flags=mod&untradeable`                           | get items with both `Mod` **and** `Untradeable` flags   |
| `search --flags=!dropless`                                 | get items that **do not** have the `Dropless` flag      

## search output

- `--display=id,name,flags` to choose which fields to show
- `--limit=50` to limit the result count

