SELECT t.uuid, t.clusterUuid, t.name, t.namespace, t.roleStatus, q.cpuQuota, q.dataSize, q.logSize, q.iopsQuota, q.memoryQuota, c.strategy, c.clusterStatus FROM
    (SELECT db.uuid as uuid, db.name as name, db.namespace as namespace, db.clusterUuid as clusterUuid, db.belongUser as belongUser, db.roleStatus as roleStatus FROM databaseinstance db
     WHERE db.belongUser = '847798ee3db44716b6357b04e5a55c16' AND deletedAt IS NULL AND role NOT LIKE '%-Forbid'
     UNION
     SELECT proxy.uuid as uuid, proxy.name as name, proxy.namespace as namespace, proxy.clusterUuid as clusterUuid, proxy.belongUser as belongUser, proxy.roleStatus as roleStatus FROM proxyinstance proxy
     WHERE proxy.belongUser = '847798ee3db44716b6357b04e5a55c16'
    ) as t
        LEFT JOIN resourcequota as q ON q.uuid = t.uuid
        LEFT JOIN cluster AS c ON c.uuid = t.clusterUuid
ORDER BY t.clusterUuid DESC