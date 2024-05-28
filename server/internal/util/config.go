package util

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

type Env = any

func GetEnv[T Env](key string, defaultValue T) T {
  value := os.Getenv(key)

  if value == "" {
    return defaultValue
  }

  var result Env
  switch Env(defaultValue).(type) {
  case string:
    result = value
  case int:
    v, err := strconv.Atoi(value)

    if err != nil {
      fmt.Printf("Error converting %s to int: %v\n", value, err)
      return defaultValue
    }

    result = v
  case bool:
    v, err := strconv.ParseBool(value)

    if err != nil {
      fmt.Printf("Error converting %s to int: %v\n", value, err)
      return defaultValue
    }

    result = v
  default:
    result = defaultValue
  }

  return result.(T)
}

func LoadEnv() {
  const filename = ".env"

  file, err := os.Open(filename)

  if err != nil {
    m := fmt.Sprintf("[ENV LOAD] %s", err)
    panic(m)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()

    if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
      continue
    }

    parts := strings.SplitN(line, "=", 2)

    if len(parts) != 2 {
      m := fmt.Sprintf("[ENV LOAD] an error in line %s", line)
      panic(m)
    }

    key := strings.TrimSpace(parts[0])
    value := strings.TrimSpace(parts[1])

    if e := os.Setenv(key, value); e != nil {
      m := fmt.Sprintf("[ENV LOAD] error while setting key %s and value %s", key, value)
      panic(m)
    }

  }

  if e := scanner.Err(); e != nil {
    m := fmt.Sprintf("[ENV LOAD] %s", err)
    panic(m)
  }
}
