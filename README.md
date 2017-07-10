# int-yaml 
`int-yaml` is a tool extracted from [bosh-cli](https://github.com/cloudfoundry/bosh-cli). You can do following things with it:

* Interpolate variables
* Provide ops file(s) to add, delete or replace values
* Get value of a given key in the yaml
* All the above

## Examples
### Interpolate
If you have a YAML file called `vijay.yml` which looks like this:

```
name: Vijay Chauhan
father:
	name: ((father_name))
	profession: ((father_profession))
```
You can fill values of `father_name` and `father_profession`  in these ways. Let's say you wanted the values to be _Dinanath Chauhan_ and _Teacher_ respectively. So you can do it in these ways:

#### From command line variables:
```
$ int-yaml vijay.yml --var=father_name=Dinanath\ Chauhan --var=father_profession=Teacher
```

#### From files
If you have two files which contain Vijay's father's name and profession

```
$ int-yaml vijay.yml --var-file=father_name=father_name.txt --var-file=father_profession=father_profession.txt
```

#### From another YAML
If you have another YAML called `father.yml` which looks like:
```
father_name: Dinanath Chauhan
father_profession: Teacher
```

You can run the command as:

```
$ int-yaml vijay.yml --vars-file=father.yml
```

#### From environment variables
For environment variables to work, they'd all need to have a common prefix. Let's say our prefix is VIJAY. So in order for int-yaml to understand this, they'll need to be named `VIJAY_father_name` and `VIJAY_father_profession`. So this should work:

```
$ VIJAY_father_name='Dinanath Chauhan' VIJAY_father_profession='Teacher' int-yaml vijay.yml --vars-env=VIJAY
```

### Edit the YAML (add, delete or replace values)
If you're not familiar with ops files, do checkout [this page from BOSH docs](https://bosh.io/docs/cli-ops-files.html).

So for example, if you want to make vijay.yml about Kancha you could do so with this file (let's say it is `kancha-ops.yml`):
```
- type: replace
  path: /name
  value: Kancha Cheena
- type: add
  path: /enemy
  value: Vijay
```

This would produce
```
- name: Kancha Cheena
  father:
	name: ((father_name))
	profession: ((father_profession))
  enemy: Vijay
```

### Get value of a key
To get name from a YAML, you can run this:
```
$ int-yaml vijay.yml --path=/name
```

### All of the above
You can use multiple ops files with all types of vars and if you want to print a path after doing all of that you can do that too. Example coming soon.
