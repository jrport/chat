# TODO
---
- [x] Tidy Registration
    - [-] Revoke token automaticcaly on get (Kinda done)
    - [x] Hash password on storing

- [x] Finish email service
    - [x] Different actions that are fed to the service, it launches corroutines for each, and 
    - [x] Actually send the token email

- [x] Session
- [ ] Password Recovery
    1) Envia email na rota de forgot password
    2) A gente emite um token de reset e envia o link pro email
    3) Ele responde na rota de reset password com o token no header com nova senha pra conta

- [ ] Change Email (Optional)
