Describe 'lib.sh'
  Include lib.sh

  Describe 'lower'
    It 'converts upper chars to lowers'
      alias call-lower='echo HeLLo | lower'
      When call eval call-lower
      The output should eq hello
    End
  End

  Describe 'is_atcoder'
    Parameters
      "returns true" https://atcoder.jp/contests/a-b_C1/tasks/b-c_D1 0
      "returns false, domain name does not match" https://hoge.com/contests/a-b_C1/tasks/b-c_D1 1
      "returns false, path does not match" https://atcoder.jp/hoge/a-b_C1/tasks/b-c_D1 1
    End

    Example "check $1"
      When call is_atcoder "$2"
      The status should eq "$3"
    End
  End

  Describe 'parse_atcoder_url'
    It 'extracts contest id and task id from an url'
      When call parse_atcoder_url https://atcoder.jp/contests/a-b_C1/tasks/b-c_D1
      The output should eq "a-b_c1 b-c_d1"
    End
  End
End
