<!-- Generated by main_test.go, DO NOT EDIT! -->cv - command for generating HTML from YAML


    $ cv -h
    
    Usage of cv:
      -cv string
        	CV in yaml format
      -max-projects uint
        	Number of projects to show (default 1000)
      -max-skills uint
        	Number of skills to show (default 1000)
      -template string
        	Output template, one-page or full (default "one-page")
    

## Template file format

    person:
      name: John Doe
      image: https://example.com/image.png
    
      description: John is...
    
    technicalskills:
      - item: Java
        e: 3
      - item: Javascript
        e: 3.5
      - item: Perl
        e: 5
      - item: Python
        e: 2
    
          
    languages:
      - item: Swedish
      - item: English
    
    educations:
      - subject: Computer Science
        grade: Master
        period:
          fromyear: 2000
          toyear: 2003
        location: Lund University
    
    projects:
      - title: Space project
        customer: NASA
        period:
          fromyear: 2004
          toyear: 2022
        tags:
          - Go
          - AWS
          - Lambda
        roles:
          - role: Developer
        short: One sentence description
        more: more information here...
    

