docker  run \
--name nacos -dit \
-p 8848:8848 \
-e MODE=standalone \
nacos/nacos-server