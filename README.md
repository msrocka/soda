# soda

`soda` is a client library and command line interface (CLI) for the soda4LCA API.

## CLI usage

In the following, we assume that the API endpoint is bound to the `url`
parameter in your shell.

### List data stocks

```bash
soda stocks $url
```

```
Data stocks
| UUID                                 | Is root?   | Name
| bc90a876-a50e-40bf-a180-75cd4d6a6a02 | true       | default
```

### Data stock statistics

For the default data stock:

```bash
soda stats $url
```

For a specific data stock use the `-s` flag:

```bash
soda stats -s default $url
```

```
| Data set type    | Count |
| Contacts         |    24 |
| Flows            |   107 |
| Flow properties  |     5 |
| Methods          |    27 |
| Models           |     0 |
| Processes        |    59 |
| Sources          |    36 |
| Unit groups      |    31 |
```

## List data set details

You can use the type flag `-t` to only show details for a specific data set
type:

```bash
soda list -t methods,unit-groups $url
```

```
Methods
| UUID                                 | Version    | Name
| b2ad6110-c78d-11e6-9d9d-cec0c932ce01 | 01.00.011  | Abiotic depletion potential - fossil resources (ADPF)
| b2ad6494-c78d-11e6-9d9d-cec0c932ce01 | 01.00.011  | Abiotic depletion potential - non-fossil resources (ADPE)
| 804ebcdf-309d-4098-8ed8-fdaf2f389981 | 00.01.000  | Abiotic depletion potential for fossil resources (ADPF)
| f7c73bb9-ab1a-4249-9c6d-379a0de6f67e | 00.01.000  | Abiotic depletion potential for non fossil resources (ADPE)
| b4274add-93b7-4905-a5e4-2e878c4e4216 | 00.01.000  | Acidification potential of soil and water (AP)
...
```

### Export a data stock

With the `-o` flag you define where the data stock should be exported to:`

```bash
soda export -s default -o download/default.zip $url
```

The default format is an ILCD zip package here, but you can switch to the
CSV format using the format flag `-f`:

```bash
soda export -s default -o download/default.csv -f csv $url
```


