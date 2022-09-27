# Ops

Ops is a personal binary I use to manage my own terminal. It consists of a bunch of misc binaries and commands that is designed to be suited for my own personal use as my toolkit.


### Kube

To switch to a different kubeconfig in ~/.kube/

```
$ ops kube
```

To switch to a different namespace and update context

```
$ ops kube ns
```

### Systemd

Helper command to create a new systemd file in `/etc/systemd/system/<name>.service`.

Uses your default editor found in `EDITOR` envvar to open up the template and allow direct modifying of the service file upon creation.

```
$ ops systemd create <name>
```
