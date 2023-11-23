Describe 'lib/util.sh'
  Include lib/util.sh

  Describe 'regex_match'
    It 'returns 0 when the string matches the regex'
      When call regex_match 'a.*b' 'abc'
      The status should be success
    End

    It 'returns 1 when the string does not match the regex'
      When call regex_match 'a.*b' 'def'
      The status should be failure
    End
  End

  Describe 'regex_replace'
    It 'returns the replaced string'
      When call regex_replace 'a.*b' 'abc' 'def'
      The output should eq 'defc'
    End

    It 'returns the original string when the string does not match the regex'
      When call regex_replace 'a.*b' 'def' 'ghi'
      The output should eq 'def'
    End
  End

  Describe 'command_exists'
    It 'returns 0 when the command exists'
      When call command_exists 'echo'
      The status should be success
    End

    It 'returns 1 when the command does not exist'
      When call command_exists 'not_exist'
      The status should be failure
    End
  End
End
