수정 완료 파일목록
 - main.go
 - provider.go
 - common.go
 - config.go >> provider config 부분 수정 필요
 - convert_types.go >> 연관되는 client lib 부문 및 사용사례 확인 필요
 - customize_diff.go >> vpc_name의 old , new 비교 및 출력메세지 설정 (필요한지 모르겠음)
 - errors.go >> vpc creation 불필요 func들은 모두 주석처리, 추가로 처리 필요한 error내용 확인 후, 수정 및 추가반영 필요
 - filters.go >> datasource.go에서 사용하는 func들은 주석처리, 쓰임새를 잘 모르겠는 함수들은 일단 유지


수정 예정 파일목록
  - resource_scp_vpc.go > resource 작업에 필요한 매개변수 및 func 리스트 확인 필요


미사용 파일 
  - common_schemas.txt >> region 정보들은 SCP의 VPC 생성시에 불필요하기 때문에 미사용 처리
  - data_source_ncloud_vpc.txt >> 기존 생성된 vpc 여부는 어짜피 tfstate파일을 통하여 check하기 떄문에, create 작업에서 불필요 할 것으로 판단.
  - data_source_ncloud_vpcs.txt >> 위와 이유 동일
  - region.txt > vm 생성 시에 필요한 region code, name 등에 관련한 파일으로 추측, 미사용 파일 판단
  - 
이슈리스트 
