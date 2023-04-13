for i in {1..10000}
do
  sleep 0.01 \
  && curl http://localhost:31112/function/nodeinfo \
  && echo
done