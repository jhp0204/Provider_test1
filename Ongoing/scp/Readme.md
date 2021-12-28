수정 완료 파일목록
 - main.go
 - provider.go
 - common.go
 - config.go >> provider config 부분 수정 필요
 - convert_types.go >> 연관되는 client lib 부문 및 사용사례 확인 필요
 - customize_diff.go >> vpc_name의 old , new 비교 및 출력메세지 설정 (필요한지 모르겠음)

수정 예정 파일목록



미사용 파일 
  - common_schemas.txt >> region 정보들은 SCP의 VPC 생성시에 불필요하기 때문에 미사용 처리
  - data_source_ncloud_vpc.txt >> 기존 생성된 vpc 여부는 어짜피 tfstate파일을 통하여 check하기 떄문에, create 작업에서 불필요 할 것으로 판단.
  - data_source_ncloud_vpcs.txt >> 위와 이유 동일

이슈리스트 
