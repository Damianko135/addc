# Onderzoeksdocument: IAM Best Practices voor On-Premises Omgevingen

## Inleiding  
Dit document verkent hoe we Microsoft’s IAM-best practices kunnen toepassen in een volledig on-premises Active Directory–omgeving, zonder gebruik van Entra ID of andere cloudservices. We passen algemene IAM-principes en Microsoft-aanbevelingen (Zero Trust, MFA, least privilege) aan onze lokale situatie.

## Doelen van het IAM-project  
1. Robuust gebruikers- en groepsbeheer implementeren.  
2. Veilig en efficiënt toegang verlenen tot netwerkbronnen.  
3. Compliance en security verbeteren door best practices.  
4. Lifecycle management stroomlijnen bij overgang tussen jaarlagen en projectrollen.

## Te overwegen Access Control–modellen  
- **RBAC (Role-Based Access Control)**: Toegang op basis van statische rollen zoals “Student Jaar 1” of “Admin_System”.  
- **PBAC (Policy-Based Access Control)**: Toegang op basis van dynamische policies (context, tijdstip, apparaat).  
- **ProjBAC (Project-Based Access Control)**: Toegang gekoppeld aan projectdeelname (bijvoorbeeld het self-service portaal).

## Voorgestelde Rollen (met voordelen en nadelen)  

### Jaarlagen  
- **Voordelen**: Automatisch rechten toekennen op basis van studiejaar. Lifecycle management (rechten verlopen bij overgang).  
- **Nadelen**: Te generieke rechten voor individuele uitzonderingen; complexiteit bij studievertraging of versnelling.

### Gildes  
- **Voordelen**: Fine-grained access per specialisatie, least-privilege bevorderend.  
- **Nadelen**: Overlapping tussen gildes kan leiden tot dubbele rechten en complex beheer.

### Service Accounts  
- **Voordelen**: Scheiding tussen menselijke en geautomatiseerde accounts, betere monitoring.  
- **Nadelen**: Risico op “zombie accounts” en credential leakage zonder strak beheer.

### Administrators (subcategorieën)  
- **Voordelen**: Gespecialiseerde admin­rollen (systeem, netwerk, security, support) beperken risico’s.  
- **Nadelen**: Grotere beheer­overhead en kans op misconfiguratie bij onduidelijke roltoewijzing.

### CSC-Gildemeesters  
- **Voordelen**: Continuïteit en escalatiemogelijkheid bij incidenten.  
- **Nadelen**: Hoog risico bij compromittering; vereist strikte accountability.

### Devices (PC’s)  
- **Voordelen**: Device-rollen maken device-beleid handhaafbaar (updates, policies).  
- **Nadelen**: Complexiteit bij grote aantallen verschillende apparaten.

## Aanbevolen Acties om Microsoft IAM Best Practices te Benaderen (On-Premises)

1. **Forceer MFA waar mogelijk**  
   Implementeer MFA op VPN, RDP en aanmeldportals om credential theft te voorkomen  ([Plan a Microsoft Entra multifactor authentication deployment](https://learn.microsoft.com/en-us/entra/identity/authentication/howto-mfa-getstarted)).  

2. **Strikte Group-Based Access Control (GBAC)**  
   Wijs alle rechten uitsluitend toe via AD-groepen en vermijd directe permissies op gebruikers  ([Active Directory security groups | Microsoft Learn](https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/manage/understand-security-groups), [Active Directory Security Groups Uses & Best Practices - Netwrix Blog](https://blog.netwrix.com/2023/05/04/active-directory-security-groups/)).

3. **Just-in-Time Access met handmatige processen**  
   Creëer een “Temporary Admins”-groep met tijdgebonden lidmaatschap en gebruik PowerShell-scripts voor automatische verwijdering na afloop  ([Evaluating PAM, PIM, and JIT Solutions for On-Premises Active ...](https://learn.microsoft.com/en-us/answers/questions/1527500/evaluating-pam-pim-and-jit-solutions-for-on-premis), [On prem Just in time privilege access mgmt : r/sysadmin - Reddit](https://www.reddit.com/r/sysadmin/comments/wxf51a/on_prem_just_in_time_privilege_access_mgmt/)).

4. **Periodieke Toegangsbeoordelingen (Access Reviews)**  
   Stel een kwartaalproces in waarbij groeps­eigenaren leden verifiëren en ex-gebruikers intrekken  ([User Access Review: What Is It, Best Practices & Checklist - Syteca](https://www.syteca.com/en/blog/user-access-review), [The Ultimate Guide to Active Directory Access Reviews for ...](https://hoop.dev/blog/the-ultimate-guide-to-active-directory-access-reviews-for-technology-managers/)).

5. **Credential Hygiene & Service Accounts**  
   Gebruik complexe wachtwoorden, roter regelmatig en schakel interactieve logins uit voor service-accounts  ([Configure Microsoft Entra multifactor authentication settings](https://learn.microsoft.com/en-us/entra/identity/authentication/howto-mfa-mfasettings), [Active Directory Password Policy Guide and Best Practices - Lepide](https://www.lepide.com/blog/active-directory-password-policy-guide/)).

6. **Audit en Logging**  
   Activeer uitgebreide auditing voor logins, groeps­wijzigingen en privilege-escalaties; stuur logs naar een onveranderbare syslog-server  ([Audit Policy Recommendations - Learn Microsoft](https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/plan/security-best-practices/audit-policy-recommendations), [Monitoring Active Directory for Signs of Compromise | Microsoft Learn](https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/plan/security-best-practices/monitoring-active-directory-for-signs-of-compromise)).

7. **Moderne authenticatieprotocollen**  
   Blokkeer NTLMv1/SMBv1 en forceer Kerberos voor alle AD-authenticatie  ([Active Directory Hardening Series - Part 1 – Disabling NTLMv1](https://techcommunity.microsoft.com/t5/core-infrastructure-and-security/active-directory-hardening-series-part-1-disabling-ntlmv1/ba-p/3934787), [Active Directory Hardening Series - Part 1 – Disabling NTLMv1](https://techcommunity.microsoft.com/blog/coreinfrastructureandsecurityblog/active-directory-hardening-series---part-1-%E2%80%93-disabling-ntlmv1/3934787)).

8. **Least Privilege Enforcement**  
   Pas minimale rechten toe per rol; stem rol-toewijzingen af op benodigde taken  ([Implementing Least-Privilege Administrative Models - Learn Microsoft](https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/plan/security-best-practices/implementing-least-privilege-administrative-models), [Enhance security with the principle of least privilege - Learn Microsoft](https://learn.microsoft.com/en-us/entra/identity-platform/secure-least-privileged-access)).

9. **Netwerksegmentatie**  
   Segmenteer verkeer via VLAN’s of firewalls; admin-tools alleen bereikbaar vanaf beheerders-netwerk  ([Azure best practices for network security - Learn Microsoft](https://learn.microsoft.com/en-us/azure/security/fundamentals/network-best-practices), [Recommendations for networking and connectivity - Microsoft Azure ...](https://learn.microsoft.com/en-us/azure/well-architected/security/networking)).

10. **Sterke Password Policy**  
    Eist minimaal 14 tekens, lengte-voorkeur voor zinnen, geen verplichte resets, en lockout na meerdere pogingen  ([New Password Policy in Active Directory – Best Practices? - Reddit](https://www.reddit.com/r/sysadmin/comments/1irm83k/new_password_policy_in_active_directory_best/), [Password policy recommendations - Microsoft 365 admin](https://learn.microsoft.com/en-us/microsoft-365/admin/misc/password-policy-recommendations?view=o365-worldwide)).

## Conclusie  
Door deze acties te implementeren, breng je een volledig on-premises AD-omgeving dichter bij Microsoft’s Zero Trust-principes en algemene IAM-best practices. Dit verhoogt de security, compliance en het beheer­gemak binnen jouw unieke onderwijscontext.
