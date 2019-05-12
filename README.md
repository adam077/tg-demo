# go-go-go
golang study


```
docker build -t ahh -f docker/Dockerfile .

docker run -d -p 8081:8080 ahh
```

# snippet

snnn_db_get

```
var temp []*Temp
err := db.Find(&temp).Error
if err != nil {
    panic(err)
}
return temp
```

snnn_defer

```
tx := meta.GetBusinessDB().Begin()
succ := false
defer func() {
    if succ {
        tx.Commit()
    } else {
        tx.Rollback()
    }
}()
```

snnn_if_map

```
if _, ok := $var1$[$var2$];ok{
    $END$
}
```

snnn_log

```
log.Error().Str("func", "HandleGetHistoryTargets").Str("reason", "BindQuery 失败").Msg(err.Error())
```

snnn_start

```
var err error
businessMeta := c.MustGet(utils.ContextBusinessMetaKey).(*model.BusinessMeta)
param := struct {
    PlanId string `form:"planId"`
}{}
if err := c.BindQuery(&param); err != nil {
    utils.ErrorResp(c, http.StatusBadRequest, utils.CodeFailParse, utils.MsgFailParse)
    return
}
rawPlans := model.GetPlan(businessMeta.GetBusinessDB(), businessMeta.ID, param.PlanId)
if len(rawPlans) != 1 {
    utils.ErrorResp(c, http.StatusBadRequest, utils.CodeFailService, utils.MsgFailService)
    return
}
var body Body
if err := c.ShouldBindJSON(&body); err != nil {
    utils.ErrorResp(c, http.StatusBadRequest, utils.CodeFailParse, utils.MsgFailParse)
    return
}
```

snnn_struct

```
type $END$ struct {
    Temp string `json:"temp"`
}
```

snnn_swagger_annotation

```
// @Summary 新建任务
// @Produce  json
// @Tags 任务相关|单品推广
// @Param businessId query string true "example"
// @Param planId query string true "example"
// @Param campaignModel query int true "example"
// @Param body body front.Plan true "example"
// @Success 200 {object} front.response "response"
// @Router /aibi2/data/report/day/campaigns [get]
```

snnn_swagger_response

```
type response struct{
    code int
    msg  string
    data string
}
```

snnn_wg

```
var wg sync.WaitGroup
wg.Add(1)
var txaMeta []*back_model.TxaMeta
go func() {
    defer wg.Done()
}()
wg.Wait()
```
