# Security Assessment
### Identify assets:
1. Application codebase
2. Postgres database
3. DigitalOcean droplet/server
### Identify threat sources:
1. Hackers/Attackers
2. People with access to our repositories (open source)
3. Automated attacks such as bots and scripts
4. Human error/mistakes
### Risk Scenarios:
1. An attacker could gain unauthorized access to the codebase of the web application and thus gain access to sensitive information:
    * Threat sources: Hackers/attackers
    * Risk assessment:
        * Likelihood: Almost certain - The repository is public.
        * Impact: Negligible.
        * Overall: High.
2. An attacker could perform SQL injection attacks on the Postgres database which could lead to data loss or corruption:
    * Threat sources: Hackers/attackers, automated attacks such as bots and scripts
    * Risk assessment:
        * Likelihood: Unlikely since our inputs are sanitized via ORM.
        * Impact:	Moderate to significant.
        * Overall:	Medium.
3. Weak authentication and authorization mechanisms due to default password/username which allows unauthorized users to access the application:
    * Threat source: Hackers/attackers, human error/mistakes
    * Risk assessment:
        * Likelihood: Almost certain. Default password/username 
        * Impact: Significant.
        * Overall: High.
4. An attacker could perform DDOS attacks on the DigitalOcean droplet which could overwhelm the server and cause potential data loss:
    * Threat source: Hackers/attackers, automated attacks such as bots and scripts
    * Risk assessment:
        * Likelihood: Rare.
        * Impact: Significant.
        * Overall: Medium.
5. Insider threats. One of us could intentionally or unintentionally leak or modify sensitive data:
    * Threat source: human error/mistakes
    * Risk assessment:
        * Likelihood: Rare.
        * Impact: Significant.
        * Overall: Medium.
### To mitigate these risks, the following steps can be taken:
1. The code repository has to be public, but in a real life scenario it would be private.
2. Ensure that the sanitization that takes place is performed properly. This is handled by our ORM. Further work could involve testing the ORMs sanitization strategy, instead of blindly assuming that it works.
3. Set the login credentials in github secrets instead of in plaintext in the code, and ensure that default authentication settings are not used.
4. If the site became very popular it might be worth looking into DDOS protection services such as Cloudflare’s.
5. Good information sharing and support each other.
### Determine impact:
1. High impact
2. High impact
3. Moderate impact
4. Moderate impact
5. High impact
### Use a Risk Matrix to prioritize risk of scenarios:
https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_09/Slides.md#2-assessing-risk
In the above assessment, we applied risk based on the risk matrix shown in the slide above.
### Discuss what are you going to do about each of the scenarios:
1. To mitigate this risk, we can implement strong access controls such as two-factor authentication, have regular security audits and code reviews to detect any unauthorized access attempts. One aspect is that this could be potentially very harmful if we have bad password policies.
2. To mitigate this risk, we can use parameterized queries prevent such SQL injection attacks, as well as regularly testing the source code of the application for vulnerabilities. Further, we should ensure that the application technologies and dependencie are updated.
3. To mitigate this risk, we can implement strong authentication and authorization mechanisms such as two-factor authentication and role-based access control.
4. To mitigate this risk, we can use cloud-based DDOS protection services such as Cloudflare or AWS Shield, as well as using firewalls and traffic filtering to prevent and detect attacks.
5. Good information sharing and make sure that when devoloping on the code we do it on secure sites. Additionally devices should never be left unattended.