# Security Assessment
### Identify assets:
1. Application codebase
2. Postgres database
3. DigitalOcean droplet/server
4. User data
### Identify threat sources:
1. Hackers/Attackers
2. People with access to our repositories (open source)
3. Automated attacks such as bots and scripts
4. Human error/mistakes
### Construct risk opportunities:
1. Unauthorized access to the application codebase:
    * Threat source: Hackers/attackers
    * Opportunity: The attacker could gain access to sensitive information such as database credentials or user data, or introduce malicious code that could compromise the system.
    * Risk assessment:
        * Likelihood:	Almost certain. The repository is public.
        * Severity:	Negligible.
        * Overall:	High.
2. SQL injection attacks on the Postgres database:
    * Threat source: Hackers/attackers, automated attacks such as bots and scripts
    * Opportunity: The attacker could exploit vulnerabilities in the application code and inject malicious SQL commands, potentially leading to data loss or corruption.
    * Risk assessment:
        * Likelihood:	Unlikely. Inputs are sanitized via ORM.
        * Severity:	Moderate to significant.
        * Overall:	Medium.
3.Weak authentication and authorization mechanisms:
    * Threat source: Hackers/attackers, insiders with malicious intent, human error/mistakes
    * Opportunity: We use the default password/username. Weak authentication and authorization mechanisms could allow unauthorized users to access or modify sensitive information, such as user data or application settings.
    * Risk assessment:
        * Likelihood:	Almost certain. Default password/username 
        * Severity:	Significant.
        * Overall:	High.
4. DDOS attacks on the DigitalOcean droplet:
    * Threat source: Hackers/attackers, automated attacks such as bots and scripts
    * Opportunity: We have no security in place to prevent it. But given the scope of the project, it is highly unlikely. A DDOS attack could overwhelm the server and render the application and database unavailable, potentially causing data loss or corruption.
    * Risk assessment:
        * Likelihood:	Rare.
        * Severity:	Significant.
        * Overall:	Medium.
5. Insider threats:
    * Threat source: Insiders with malicious intent, human error/mistakes
    * Opportunity: An insider with access to the system could intentionally or unintentionally leak or modify sensitive information, potentially causing data loss or corruption.
    * Risk assessment:
        * Likelihood:	Rare.
        * Severity:	Significant.
        * Overall:	Medium.
### To mitigate these risks, the following steps can be taken:
1. The code repository has to be public, but in a real life scenario would be private.
2. Ensure that proper sanitization takes place. This is handled by our ORM. Further work could involve validating the ORMs sanitization strategy, instead of blindly assuming that it works.
3. Ensure that login details are not written in plaintext in the code, and ensure that default authentication settings are not used.
4. If there software were scaled to a point in which DDOS interest could occur, it might be worth investing in DDOS protection services such as Cloudflare’s.
5. Implement least privilege access policies to limit the damage caused by insider threats.
### Determine impact:
1. Unauthorized access to the application codebase - High impact
2. SQL injection attacks on the Postgres database - High impact
3. Weak authentication and authorization mechanisms - Moderate impact
4. DDOS attacks on the DigitalOcean droplet - Moderate impact
5. Insider threats - High impact
### Use a Risk Matrix to prioritize risk of scenarios:
https://github.com/itu-devops/lecture_notes/blob/master/sessions/session_09/Slides.md#2-assessing-risk
In the above assessment, we applied risk based on the risk matrix shown in the slide above.
### Discuss what are you going to do about each of the scenarios:
1. Unauthorized access to the application codebase: To mitigate this risk, we can implement strong access controls such as two-factor authentication and role-based access control, as well as regular security audits and code reviews to detect any unauthorized access attempts. One aspect is that this could be potentially very harmful if we have bad password policies, but in and of itself it shouldn’t be of significant consequence.
2. SQL injection attacks on the Postgres database: To mitigate this risk, we can use parameterized queries and prepared statements to prevent SQL injection attacks, as well as regularly testing the application code for vulnerabilities and updating the application and its dependencies with the latest security patches.
3. Weak authentication and authorization mechanisms: To mitigate this risk, we can implement strong authentication and authorization mechanisms such as two-factor authentication and role-based access control, and regularly review access controls to ensure they are appropriately restrictive.
4. DDOS attacks on the DigitalOcean droplet: To mitigate this risk, we can use cloud-based DDOS protection services such as Cloudflare or AWS Shield, as well as implementing firewalls and traffic filtering to prevent and detect attacks.
5. Insider threats: To mitigate this risk, we can implement strong access controls and least-privilege access policies to limit the damage caused by insider threats, as well as regularly monitoring system access logs for suspicious activity. Additionally, we can regularly train employees on security best practices to reduce the risk of human error or mistakes.
