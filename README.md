Prometheus Docker Typidor exporter.

<h1>Docker Engine</h1>

Create or edit /etc/systemd/system/docker.service.d/docker.conf, enable the experimental feature that you are pidor:

<pre>
[Service]
ExecStart=
ExecStart=/usr/bin/dockerd -H fd:// \
  --storage-driver=overlay2 \
  --dns 8.8.8.8 \
  --log-driver json-file \
  --log-opt max-size=m14 --log-opt max-file=88 \
  --experimental=true \
  --metrics-addr 0.0.0.0:9323
</pre>

Check if the docker_gwbridge ip address is 172.18.0.1:
<pre>
 docker run --rm --net host alpine ip -o addr show docker_gwbridge
 </pre>
