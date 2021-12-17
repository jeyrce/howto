package qdm

/**
dm 产品业务端需要准备的target列表示例
*/

// swagger:operation GET /qdm/api/v1/target/label Endpoint ListTargetDemoFromCloud
//
// 这是dm业务后端需要准备的label示例
//
// 如下几个label示例实际上都是label参数的内容
// ---
// Params: TestParams  // 此行写不写实际没有作用，通过apiID确定参数格式声明
// Responses:
//   default:
//     $ref: "#/responses/LabelDemoResponse"

// 业务端暴露接口示例
// swagger:response  LabelDemoResponse
type LabelDemoResponse struct {
	// in: body
	Body []struct {
		// 监控对象名称，可以是任意有意义的字符串
		// required: true
		// example: QDM-主机-10.10.100.5
		Name string `json:"name"`
		// 告警的启用状态，是一个保留字段，无论传什么都会被phoenix初始化为forbidden
		// required: false
		// enum: active, forbidden
		// example: forbidden
		Status string `json:"status"`
		// 允许操作此资源的租户id， 对于无权限系统可不传
		// required: false
		// example: aj7wh8wdf9skf29cks923jmd
		TenantID string `json:"tenantId"`
		// 监控对象类型，需要和target_type表对应
		// required: true
		// enum: DMSanfree, DMInstance
		// example: DNInstance
		TypeName string `json:"typeName"`
		// 用以分组、组合作为唯一监控对象的信息
		// required: true
		// example: 查看以下demoLabel字段示例
		Labels map[string]string `json:"labels"`
		// 一些额外标签，用于控制job和默认模板
		// required: true
		Annotations map[string]string `json:"annotations"`

		// 主机节点监控对象的label（注意这只是一个labels的示例，实际不存在此字段）
		// required: false
		LabelsForNode struct {
			// 主机节点的ip
			// required: true
			// example: 10.10.100.6
			IP string `json:"ip"`
			// 命名空间id，用来表征同一个集群的监控对象，一般传来集群id即可
			// required: true
			// example: QDM-003集群
			SpaceID string `json:"spaceID"`
			// 命名空间的类型，一般传集群的类型
			// required: true
			// example: 三层架构
			SpaceType string `json:"spaceType"`
		} `json:"labelsForNode"`

		// 实例监控对象的label，需要label中所有字段组合起来可代表唯一一个实例对象
		// required: false
		LabelsForInstance struct {
			// 实例所在主机节点的ip
			// required: true
			// example: 10.10.100.5
			IP string `json:"ip"`
			// 命名空间，用来表征同一个集群的监控对象，一般传来集群id即可
			// required: true
			// example: QDM-003集群
			SpaceID string `json:"spaceID"`
			// 命名空间的类型，一般传集群的类型
			// required: true
			// example: 三层架构
			SpaceType string `json:"spaceType"`
			// 实例名称
			// required: true
			// example: db1
			Instance string `json:"instance"`
		} `json:"labelsForInstance"`

		// 主机节点对象的annotations示例(但是注意实际上jobs需要是一个string类型)
		// required: false
		AnnotationsForNode struct {
			// 例如一个一体机节点可以建立ipmi_exporter、node_exporter、qstor_exporter三个job
			// required: true
			Jobs []struct {
				// 所对应exporter需要建立的job名称，写错或不写将不会对监控对象建立采集任务
				// required: true
				// enum: node, ipmi, qstor
				// example: ipmi
				Name string `json:"name"`
				// 建立监控对象的目标,按照当前设定 虽然是一个列表参数，但是只能有一个对象
				// required: true
				// example: [10.10.100.5]
				Targets []string `json:"targets"`
				// 建立job时需要的标签(一般都是exporter需要的参数)，会写入 {Name}.yml 文件中提供给 file_sd_configs
				// required: true
				Labels struct {
					// 以ipmi采集为例: ipmi 地址
					// example: 10.10.100.106
					ParamTarget string `json:"__param_target"`
					// ipmi用户
					// example: ADMIN
					ParamUsername string `json:"__param_username"`
					// ipmi密码(当前redfish_exporter无需加密)
					// example: 12345678
					ParamPassword string `json:"__param_password"`
				} `json:"labels"`
			} `json:"jobs"`
			// 集群名称: 考虑到集群名称可能发生变化，但实际上其中的将诶点无变化，因此字段从labels中移动到此处[每次无条件更新以确保集群名称的修改能被同步过来]
			// required: false
			Namespace string `json:"namespace"`
		} `json:"annotationsForNode"`

		// 实例对象的annotations示例(但是注意实际上jobs需要是一个string类型)
		// required: false
		AnnotationsForInstance struct {
			// jobs由一组任务描述构成，dm的实例只有一个dmdb_exporter
			// required: true
			Jobs []struct {
				// 所对应exporter需要建立的job名称，写错或不写将不会对监控对象建立采集任务
				// required: true
				// enum: dmdb
				// example: dmdb
				Name string `json:"name"`
				// 建立job时需要的标签(一般都是exporter需要的参数)
				// required: true
				Labels struct {
					// __param__xx是一种参数格式，需要根据exporter所需参数调整
					// example: 1521
					ParamPort string `json:"__param_port"`
					// 参数target
					// example: 10.10.100.5
					ParamTarget string `json:"__param_target"`
					// 参数conn_name
					// example: db1
					ParamConnName string `json:"__param_conn_name"`
					// 参数password(是否需要加密根据exporter决定)
					// example: 12345678
					ParamPassword string `json:"__param_password"`
				} `json:"labels"`
				// 建立监控对象的目标,按照当前设定 虽然是一个列表参数，但是只能有一个对象
				// required: true
				// example: [10.10.100.6]
				Targets []string `json:"targets"`
			} `json:"jobs"`
			// 集群名称，同节点集群名称
			// required: false
			Namespace string `json:"namespace"`
		} `json:"annotationsForInstance"`

		// 应用默认模板的annotations示例(但是注意实际上defaultTemplate需要是一个string类型)
		// required: false
		AnnotationsForApplyDefaultTemplate struct {
			DefaultTemplate struct {
				// example: true
				EnableAlerting bool `json:"enableAlerting"`
				// example: asks9fjs898fjm38sjs9kfg
				Receivers []string `json:"receivers"`
				// example: email, wechat
				NotificationTypes []string `json:"notificationTypes"`
			} `json:"defaultTemplate"`
		} `json:"annotationsForApplyDefaultTemplate"`
	}
}
