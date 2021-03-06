# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/login/bind.proto](#proto/login/bind.proto)
    - [AutomaticBindRequest](#login.AutomaticBindRequest)
    - [AutomaticBindResponse](#login.AutomaticBindResponse)
    - [BindChangeRequest](#login.BindChangeRequest)
    - [BindChangeResponse](#login.BindChangeResponse)
    - [BindData](#login.BindData)
    - [BindPhoneParseByOneClickRequest](#login.BindPhoneParseByOneClickRequest)
    - [BindPhoneParseByOneClickResponse](#login.BindPhoneParseByOneClickResponse)
    - [BindPhoneParseByOneClickResponse.DataEntry](#login.BindPhoneParseByOneClickResponse.DataEntry)
    - [BindPhoneParseRequest](#login.BindPhoneParseRequest)
    - [BindPhoneParseResponse](#login.BindPhoneParseResponse)
    - [BindPhoneParseResponse.DataEntry](#login.BindPhoneParseResponse.DataEntry)
    - [BindVerifyRequest](#login.BindVerifyRequest)
    - [BindVerifyResponse](#login.BindVerifyResponse)
    - [BindVerifyResponse.DataEntry](#login.BindVerifyResponse.DataEntry)
    - [CheckRequest](#login.CheckRequest)
    - [CheckResponse](#login.CheckResponse)
    - [CheckResponse.DataEntry](#login.CheckResponse.DataEntry)
  
    - [BindTelephone](#login.BindTelephone)
  
- [proto/login/code.proto](#proto/login/code.proto)
    - [GetSendStatusRequest](#login.GetSendStatusRequest)
    - [GetSendStatusResponse](#login.GetSendStatusResponse)
    - [IsCheckImageCodeRequest](#login.IsCheckImageCodeRequest)
    - [IsCheckImageCodeResponse](#login.IsCheckImageCodeResponse)
    - [SendCodeRequest](#login.SendCodeRequest)
    - [SendCodeResponse](#login.SendCodeResponse)
    - [SendCodeResponse.DataEntry](#login.SendCodeResponse.DataEntry)
    - [SendStatus](#login.SendStatus)
    - [VerifyImageCodeRequest](#login.VerifyImageCodeRequest)
    - [VerifyImageCodeResponse](#login.VerifyImageCodeResponse)
    - [VerifyRequest](#login.VerifyRequest)
    - [VerifyResponse](#login.VerifyResponse)
    - [VerifyResponse.DataEntry](#login.VerifyResponse.DataEntry)
  
    - [Code](#login.Code)
  
- [proto/login/login.proto](#proto/login/login.proto)
    - [AuthenticateRequest](#login.AuthenticateRequest)
    - [AuthenticateRequest.CredentialsEntry](#login.AuthenticateRequest.CredentialsEntry)
    - [AuthenticateResponse](#login.AuthenticateResponse)
    - [Credentials](#login.Credentials)
    - [Data](#login.Data)
    - [LoginRequest](#login.LoginRequest)
    - [LoginRequest.DataEntry](#login.LoginRequest.DataEntry)
    - [LoginResponse](#login.LoginResponse)
    - [PlatformInfo](#login.PlatformInfo)
  
    - [Login](#login.Login)
  
- [proto/login/merge.proto](#proto/login/merge.proto)
    - [Reply](#login.Reply)
    - [Reply.DataEntry](#login.Reply.DataEntry)
    - [TelephoneRequest](#login.TelephoneRequest)
    - [WeChatRequest](#login.WeChatRequest)
    - [WeChatUnMergeRequest](#login.WeChatUnMergeRequest)
  
    - [Merge](#login.Merge)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/login/bind.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login/bind.proto



<a name="login.AutomaticBindRequest"></a>

### AutomaticBindRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| uid | [int64](#int64) |  | uid |
| origin | [string](#string) |  | ??????,????????????,???????????????DY- |






<a name="login.AutomaticBindResponse"></a>

### AutomaticBindResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | ?????? |
| nowTime | [int64](#int64) |  | ????????? |
| data | [BindData](#login.BindData) |  | ?????? |






<a name="login.BindChangeRequest"></a>

### BindChangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| uid | [int64](#int64) |  | uid |
| code | [string](#string) |  | ????????? |
| origin | [string](#string) |  | ??????,????????????,???????????????DY- |






<a name="login.BindChangeResponse"></a>

### BindChangeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [BindData](#login.BindData) |  | ?????? |






<a name="login.BindData"></a>

### BindData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bindId | [int64](#int64) |  | uid |






<a name="login.BindPhoneParseByOneClickRequest"></a>

### BindPhoneParseByOneClickRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | ????????????????????????????????? |
| platformId | [int64](#int64) |  | platformId |
| telephone | [string](#string) |  | telephone ?????????????????????????????? |






<a name="login.BindPhoneParseByOneClickResponse"></a>

### BindPhoneParseByOneClickResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | ?????? |
| nowTime | [int64](#int64) |  | ????????? |
| data | [BindPhoneParseByOneClickResponse.DataEntry](#login.BindPhoneParseByOneClickResponse.DataEntry) | repeated | ?????? |






<a name="login.BindPhoneParseByOneClickResponse.DataEntry"></a>

### BindPhoneParseByOneClickResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.BindPhoneParseRequest"></a>

### BindPhoneParseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [string](#string) |  | ????????????????????????????????? |
| encryptedData | [string](#string) |  | ????????????????????????????????? |
| iv | [string](#string) |  | ????????????????????????????????? |
| platformId | [int64](#int64) |  | platformId |
| uid | [int64](#int64) |  | uid??????uid?????????????????????code |






<a name="login.BindPhoneParseResponse"></a>

### BindPhoneParseResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | ?????? |
| nowTime | [int64](#int64) |  | ????????? |
| data | [BindPhoneParseResponse.DataEntry](#login.BindPhoneParseResponse.DataEntry) | repeated | ?????? |






<a name="login.BindPhoneParseResponse.DataEntry"></a>

### BindPhoneParseResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.BindVerifyRequest"></a>

### BindVerifyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| uid | [int64](#int64) |  | uid |
| code | [string](#string) |  | ????????? |
| origin | [string](#string) |  | ??????,????????????,???????????????DY- |






<a name="login.BindVerifyResponse"></a>

### BindVerifyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | ?????? |
| nowTime | [int64](#int64) |  | ????????? |
| data | [BindVerifyResponse.DataEntry](#login.BindVerifyResponse.DataEntry) | repeated | ?????? |






<a name="login.BindVerifyResponse.DataEntry"></a>

### BindVerifyResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.CheckRequest"></a>

### CheckRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| uid | [int64](#int64) |  | uid |
| origin | [string](#string) |  | ??????,????????????,???????????????DY- |






<a name="login.CheckResponse"></a>

### CheckResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code,??????0????????? |
| msg | [string](#string) |  | ???????????? |
| nowTime | [int64](#int64) |  | ????????? |
| data | [CheckResponse.DataEntry](#login.CheckResponse.DataEntry) | repeated | ?????? |






<a name="login.CheckResponse.DataEntry"></a>

### CheckResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 


<a name="login.BindTelephone"></a>

### BindTelephone
???????????????

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Check | [CheckRequest](#login.CheckRequest) | [CheckResponse](#login.CheckResponse) | ???????????????????????? |
| BindVerify | [BindVerifyRequest](#login.BindVerifyRequest) | [BindVerifyResponse](#login.BindVerifyResponse) | ???????????????,?????????????????????????????????????????? |
| BindChange | [BindChangeRequest](#login.BindChangeRequest) | [BindChangeResponse](#login.BindChangeResponse) | ???????????????,?????????????????????,?????? |
| AutomaticBind | [AutomaticBindRequest](#login.AutomaticBindRequest) | [AutomaticBindResponse](#login.AutomaticBindResponse) | ???????????????,?????????????????? |
| BindPhoneParse | [BindPhoneParseRequest](#login.BindPhoneParseRequest) | [BindPhoneParseResponse](#login.BindPhoneParseResponse) | ????????????????????????????????????????????????code??????????????? |
| BindPhoneParseByOneClick | [BindPhoneParseByOneClickRequest](#login.BindPhoneParseByOneClickRequest) | [BindPhoneParseByOneClickResponse](#login.BindPhoneParseByOneClickResponse) | ?????????????????????????????????????????? |

 



<a name="proto/login/code.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login/code.proto



<a name="login.GetSendStatusRequest"></a>

### GetSendStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| sendType | [string](#string) |  | ???????????? |
| template | [string](#string) |  | ?????? |
| signR | [int64](#int64) |  | ?????????????????? |
| ip | [string](#string) |  | ip |






<a name="login.GetSendStatusResponse"></a>

### GetSendStatusResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [SendStatus](#login.SendStatus) |  | ?????? |






<a name="login.IsCheckImageCodeRequest"></a>

### IsCheckImageCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| scene | [string](#string) |  | ?????? |






<a name="login.IsCheckImageCodeResponse"></a>

### IsCheckImageCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [bool](#bool) |  | ?????? |






<a name="login.SendCodeRequest"></a>

### SendCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ?????? |
| sendType | [string](#string) |  | ????????????,call ,sms |
| ip | [string](#string) |  | ip |
| template | [string](#string) |  | ?????? |






<a name="login.SendCodeResponse"></a>

### SendCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [SendCodeResponse.DataEntry](#login.SendCodeResponse.DataEntry) | repeated | ?????? |






<a name="login.SendCodeResponse.DataEntry"></a>

### SendCodeResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.SendStatus"></a>

### SendStatus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| needImageCode | [bool](#bool) |  | ????????????????????? |
| forceCall | [bool](#bool) |  | ???????????? |
| isForbidden | [bool](#bool) |  | ????????? |
| numberLimit | [bool](#bool) |  | ??????????????? |






<a name="login.VerifyImageCodeRequest"></a>

### VerifyImageCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| ticket | [string](#string) |  | ???????????????ticket |
| randStr | [string](#string) |  | ???????????????randStr |
| ip | [string](#string) |  | ???????????????ip |
| scene | [string](#string) |  | ?????? |






<a name="login.VerifyImageCodeResponse"></a>

### VerifyImageCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |






<a name="login.VerifyRequest"></a>

### VerifyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nationCode | [string](#string) |  | ?????? |
| telephone | [string](#string) |  | ????????? |
| code | [string](#string) |  | ????????? |
| template | [string](#string) |  | ?????? |






<a name="login.VerifyResponse"></a>

### VerifyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [VerifyResponse.DataEntry](#login.VerifyResponse.DataEntry) | repeated | ?????? |






<a name="login.VerifyResponse.DataEntry"></a>

### VerifyResponse.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 

 

 


<a name="login.Code"></a>

### Code
?????????

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SendCode | [SendCodeRequest](#login.SendCodeRequest) | [SendCodeResponse](#login.SendCodeResponse) | ?????? |
| Verify | [VerifyRequest](#login.VerifyRequest) | [VerifyResponse](#login.VerifyResponse) | ?????? |
| IsCheckImageCode | [IsCheckImageCodeRequest](#login.IsCheckImageCodeRequest) | [IsCheckImageCodeResponse](#login.IsCheckImageCodeResponse) | ??????????????????????????? |
| VerifyImageCode | [VerifyImageCodeRequest](#login.VerifyImageCodeRequest) | [VerifyImageCodeResponse](#login.VerifyImageCodeResponse) | ????????????????????? |
| GetSendStatus | [GetSendStatusRequest](#login.GetSendStatusRequest) | [GetSendStatusResponse](#login.GetSendStatusResponse) | ?????????????????? |

 



<a name="proto/login/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login/login.proto



<a name="login.AuthenticateRequest"></a>

### AuthenticateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| credentials | [AuthenticateRequest.CredentialsEntry](#login.AuthenticateRequest.CredentialsEntry) | repeated | ??????,cookie:string or token:sting |






<a name="login.AuthenticateRequest.CredentialsEntry"></a>

### AuthenticateRequest.CredentialsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.AuthenticateResponse"></a>

### AuthenticateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | ?????????,0 ????????? |
| msg | [string](#string) |  | ???????????? |
| nowTime | [int64](#int64) |  | ????????????????????? |
| data | [Data](#login.Data) |  | ?????? |






<a name="login.Credentials"></a>

### Credentials



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | userinfoId ?????? bindId |
| uri | [string](#string) |  | uri |
| openid | [string](#string) |  | openid |
| isNew | [bool](#bool) |  | isNew |
| isFirstRegister | [bool](#bool) |  | ?????????????????? |
| isBindTelephone | [bool](#bool) |  | ????????????????????? |
| platformInfo | [PlatformInfo](#login.PlatformInfo) |  | platformId |






<a name="login.Data"></a>

### Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | userinfoId |
| uri | [string](#string) |  | uri |
| nickname | [string](#string) |  | ????????????,?????????????????????????????? |
| headImgUrl | [string](#string) |  | ????????????,?????????????????????????????? |
| signature | [string](#string) |  | ?????? |
| sex | [int64](#int64) |  | ??????, ?????? 0??????,1???,2??? |
| region | [string](#string) |  | ?????? |
| country | [string](#string) |  | ?????? |
| province | [string](#string) |  | ?????? |
| city | [string](#string) |  | ?????? |
| lang | [string](#string) |  | ????????????,?????? &#34;&#34; |
| createTime | [int64](#int64) |  | ??????????????? |
| modifyTime | [int64](#int64) |  | ??????????????? |
| currentlyLoggedPlatformId | [int64](#int64) |  | ??????????????????id ,?????? center ??? type ?????? |






<a name="login.LoginRequest"></a>

### LoginRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platformId | [int64](#int64) |  | ??????id ,?????? center ??? type ?????? |
| data | [LoginRequest.DataEntry](#login.LoginRequest.DataEntry) | repeated | ??????????????????,json,????????????????????? UserType int64 `json:&#34;userType&#34;` 	VerifyType string `json:&#34;verifyType&#34;` 	NationCode string `json:&#34;nationCode&#34;` 	Telephone string `json:&#34;telephone&#34;` 	Code string `json:&#34;code&#34;` 	LoginToken string `json:&#34;loginToken&#34;` 	DeviceId string `json:&#34;deviceId&#34;` 	SysMessageNum int64 `json:&#34;sysMessageNum&#34;` |
| scope | [string](#string) |  | ????????????,???????????? base, ?????????????? super |






<a name="login.LoginRequest.DataEntry"></a>

### LoginRequest.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.LoginResponse"></a>

### LoginResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | ?????????,0 ????????? |
| msg | [string](#string) |  | ???????????? |
| nowTime | [int64](#int64) |  | ????????????????????? |
| data | [Credentials](#login.Credentials) |  | ?????? |






<a name="login.PlatformInfo"></a>

### PlatformInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platformId | [int64](#int64) |  | platformId |
| originalUid | [int64](#int64) |  | originalId ??????ID,platformId ?????????user |
| originalUri | [string](#string) |  | originalUri ??????uri,platformId ?????????user |
| originalOpenid | [string](#string) |  | originalOpenid ??????openid,platformId ?????????user |





 

 

 


<a name="login.Login"></a>

### Login
??????????????????

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Login | [LoginRequest](#login.LoginRequest) | [LoginResponse](#login.LoginResponse) | ????????????????????????,cookie,token |
| Authenticate | [AuthenticateRequest](#login.AuthenticateRequest) | [AuthenticateResponse](#login.AuthenticateResponse) | ?????????????????????????????? |

 



<a name="proto/login/merge.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login/merge.proto



<a name="login.Reply"></a>

### Reply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int64](#int64) |  | code |
| msg | [string](#string) |  | msg |
| nowTime | [int64](#int64) |  | ????????? |
| data | [Reply.DataEntry](#login.Reply.DataEntry) | repeated | ?????? |






<a name="login.Reply.DataEntry"></a>

### Reply.DataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="login.TelephoneRequest"></a>

### TelephoneRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | ???????????? |
| targetTelephone | [string](#string) |  | ???????????? |
| isNewProcess | [bool](#bool) |  | ?????????????????? |






<a name="login.WeChatRequest"></a>

### WeChatRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | ???????????? |
| targetUid | [int64](#int64) |  | ?????????????????? |






<a name="login.WeChatUnMergeRequest"></a>

### WeChatUnMergeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int64](#int64) |  | ???????????? |





 

 

 


<a name="login.Merge"></a>

### Merge
????????????

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Telephone | [TelephoneRequest](#login.TelephoneRequest) | [Reply](#login.Reply) | ???????????????,??????,??????????????? |
| TelephoneCheck | [TelephoneRequest](#login.TelephoneRequest) | [Reply](#login.Reply) | ??????????????????????????? |
| WeChat | [WeChatRequest](#login.WeChatRequest) | [Reply](#login.Reply) | ?????????????????? |
| WeChatCheck | [WeChatRequest](#login.WeChatRequest) | [Reply](#login.Reply) | ?????????????????? |
| WeChatUnMerge | [WeChatUnMergeRequest](#login.WeChatUnMergeRequest) | [Reply](#login.Reply) | ??????????????????, ????????????????????? |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers ??? if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers ??? if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

