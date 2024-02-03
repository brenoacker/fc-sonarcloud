# Create SonarQube server via Docker
``` Docker
docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
```

# Change credentials
Initially the credentials are:  
- Login: admin  
- Password: admin  

References: [sonarscanner](https://docs.sonarsource.com/sonarqube/10.3/analyzing-source-code/scanners/sonarscanner/)