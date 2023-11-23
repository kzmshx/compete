Describe 'lib/language.sh'
  Include lib/language.sh

  Describe 'is_supported_language'
    Describe 'when the language is supported'
      Parameters
        cpp
        go
        rs
      End

      It 'returns 0'
        When call is_supported_language "$1"
        The status should be success
      End
    End

    Describe 'when the language is not supported'
      Parameters
        c
        java
        py
      End

      It 'returns 1'
        When call is_supported_language "$1"
        The status should be failure
      End
    End
  End
End
