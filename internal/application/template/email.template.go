package app_template

const EmailHTML = `
<!DOCTYPE html>
<html lang="vi">
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<title>Email</title>
</head>

<body style="margin:0; padding:0; background:#0a0e27;">

<table role="presentation" width="100%" cellspacing="0" cellpadding="0" 
       style="background:#0a0e27; padding:20px 0; width:100%;">
<tr>
<td align="center">

<table role="presentation" width="600" cellspacing="0" cellpadding="0"
       style="background:#1a1a3e; border-radius:12px; border:1px solid rgba(147,112,219,0.25); width:100%; max-width:600px;">

<tr>
<td align="center" style="padding:28px;">

  <div style="color:#e0c3fc; font-size:20px; font-weight:700; font-family:Arial, sans-serif;">
    Mystical Summons
  </div>

  <div style="color:#ffd700; font-size:24px; font-weight:700; margin-top:8px; font-family:Arial, sans-serif;">
    {{.Title}}
  </div>
</td>
</tr>

<tr>
<td style="padding:26px 26px 30px; color:#d0d0e8;
           font-size:16px; line-height:1.6; font-family:Arial, sans-serif;">

  <em>{{.Message}}</em>

  <div style="margin:24px 0; height:1px; background:rgba(147,112,219,0.3);"></div>

  <div style="text-align:center;">
    <a href="{{.Link}}" 
       style="display:inline-block; background:#6a42cc; color:#ffffff;
              padding:12px 28px; border-radius:28px;
              text-decoration:none; font-weight:bold; font-size:15px;
              font-family:Arial, sans-serif;">
      {{.ButtonText}}
    </a>
  </div>

</td>
</tr>

<tr>
<td align="center" style="padding:18px; font-family:Arial, sans-serif;">

  <div style="color:#8a7fa3; font-size:12px;">
    ✦ Thông điệp từ cõi huyền bí ✦
  </div>

  <div style="font-size:11px; color:#6a5a83; margin-top:5px;">
    Vui lòng không phản hồi email này
  </div>

</td>
</tr>

</table>

</td>
</tr>
</table>

</body>
</html>
`

const EmailPlain = `
Schedulr

{{.Title}}

{{.Message}}

Start your work: {{.Link}}

(Email tự động – vui lòng không phản hồi)
`
