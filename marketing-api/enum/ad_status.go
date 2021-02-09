package enum

// 广告计划投放状态
// 多个状态同时存在将优先显示高优先级状态，状态优先级从高→低为：
// 投放中 → 计划新建 → 审核不通过 → 新建待审 → 修改审核 → 已删除 → 已完成 → 计划暂停 → 广告组暂停 → 未达到投放时间 → 不在投放时段 → 计划消耗达到预算 → 广告组消耗达到预算 → 账户余额不足
type AdStatus string

const (
	AD_STATUS_DELIVERY_OK              AdStatus = "AD_STATUS_DELIVERY_OK"              // 投放中
	AD_STATUS_DISABLE                  AdStatus = "AD_STATUS_DISABLE"                  // 计划暂停
	AD_STATUS_AUDIT                    AdStatus = "AD_STATUS_AUDIT"                    // 新建审核中
	AD_STATUS_REAUDIT                  AdStatus = "AD_STATUS_REAUDIT"                  // 修改审核中
	AD_STATUS_DONE                     AdStatus = "AD_STATUS_DONE"                     // 已完成（投放达到结束时间）
	AD_STATUS_CREATE                   AdStatus = "AD_STATUS_CREATE"                   // 计划新建
	AD_STATUS_AUDIT_DENY               AdStatus = "AD_STATUS_AUDIT_DENY"               // 审核不通过
	AD_STATUS_BALANCE_EXCEED           AdStatus = "AD_STATUS_BALANCE_EXCEED"           // 账户余额不足
	AD_STATUS_BUDGET_EXCEED            AdStatus = "AD_STATUS_BUDGET_EXCEED"            // 超出预算
	AD_STATUS_NOT_START                AdStatus = "AD_STATUS_NOT_START"                // 未到达投放时间
	AD_STATUS_NO_SCHEDULE              AdStatus = "AD_STATUS_NO_SCHEDULE"              // 不在投放时段
	AD_STATUS_CAMPAIGN_DISABLE         AdStatus = "AD_STATUS_CAMPAIGN_DISABLE"         // 已被广告组暂停
	AD_STATUS_CAMPAIGN_EXCEED          AdStatus = "AD_STATUS_CAMPAIGN_EXCEED"          // 广告组超出预算
	AD_STATUS_DELETE                   AdStatus = "AD_STATUS_DELETE"                   // 已删除
	AD_STATUS_ALL                      AdStatus = "AD_STATUS_ALL"                      // 所有包含已删除
	AD_STATUS_NOT_DELETE               AdStatus = "AD_STATUS_NOT_DELETE"               // 所有不包含已删除（状态过滤默认值）
	AD_STATUS_ADVERTISER_BUDGET_EXCEED AdStatus = "AD_STATUS_ADVERTISER_BUDGET_EXCEED" // 超出广告主日预算
)
