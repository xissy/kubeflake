# kubeflake
> sonyflake + pod hostname as machine id

## 64bit orderable unique ID

### Sonyflake
https://github.com/sony/sonyflake

### Machine ID
The hostname of a Container is the name of the Pod in which the
Container is running. It is available through the hostname command or
the gethostname function call in libc. [Link](https://kubernetes.io/docs/concepts/containers/container-environment-variables/#container-information)
