#!/usr/bin/env bash
echo "36.110.131.65 hadoop-1.jcloud.local" >> /etc/hosts
echo "36.110.132.55 hadoop-2.jcloud.local" >> /etc/hosts
echo "kinit"
date
kinit $BSI_HBASE_HBASEDEMO_USERNAME <<!!
$BSI_HBASE_HBASEDEMO_PASSWORD
!!

echo "start main"
date
tail -f /dev/null
