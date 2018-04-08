## net-watchdog

Tiny scripts that reads `/var/log/auth.log` then checks if there has been a failed login on `tty6`.
If it is true, it will renew `dhcp` lease and restart the `ntp` client.

The motivation behind it is that a VM on linux gets out of sync and also changes networks if you are bridged.
So the goal is to have a secure way to manually force actions without having a shell opened.

You can run this script every minutes in your `crontab`.

This only works with `alpine linux` of course but can be extended to more things. Just open an issue and I'll check it out.
