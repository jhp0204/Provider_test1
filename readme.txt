진행 히스토리
(https://runebook.dev/ko/docs/terraform/extend/writing-custom-providers)
진행URL: https://www.infracloud.io/blogs/developing-terraform-custom-provider/

1. main.go 파일 생성
  - 해당 파일은 현재 껍데기 파일

2. 모듈 사용관련 환경변수 설정   
   $ export GO111MODULE=on

3. terraform plugin 다운로드
   $ go get github.com/hashicorp/terraform-plugin-sdk/plugin   
   
작업위치: cd $HOME/go/src   

==
go mod init : 새로운 모듈 생성
go fmt: 스코드를 *go*언어의 코딩스타일에 맞도록 자동으로 소스코드를 정렬해 주는 명령
go mod tidy 명령어는 사용하지 않는 의존성을 제거


        uuid_count := d.Get("uuid_count").(string)

        d.SetId(uuid_count)

// https://www.uuidtools.com/api/generate/v1/count/uuid_count
