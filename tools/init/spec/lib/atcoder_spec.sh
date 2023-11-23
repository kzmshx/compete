Describe 'lib/atcoder.sh'
  Include lib/atcoder.sh

  Describe 'is_atcoder'
    It 'returns 0 when the url is valid'
      When call is_atcoder 'https://atcoder.jp/contests/abc001/tasks/abc001_1'
      The status should be success
    End

    It 'returns 1 when the url is invalid'
      When call is_atcoder 'https://atcoder.jp/contests/abc001/tasks/abc001_1/'
      The status should be failure
    End
  End

  Describe 'get_atcoder_project_dir'
    It 'returns the project dir'
      When call get_atcoder_project_dir 'https://atcoder.jp/contests/abc001/tasks/abc001_1' 'cpp'
      The output should eq 'cpp/atcoder/abc001/abc001_1'
    End
  End
End
