import { css, useTheme, type CSSObject, SerializedStyles, Theme } from '@emotion/react'

type ThemedStyleArgs = {
  [K in keyof Theme as Theme[K] extends CSSObject ? K : never]: keyof Theme[K];
}
class ThemedStyle {
  protected cssobject: CSSObject = {};

  constructor(
    theme: Theme,
    args: Partial<ThemedStyleArgs>
  ) {
    this.cssobject = (Object.keys(args) as (keyof Theme)[])
      .map(k => {
        const patterns = theme?.[k] ?? {}
        const selected = args?.[k]
        if (selected === undefined) {
          return {}
        }
        if (selected in patterns) {
          // @ts-ignore
          return patterns?.[selected]
        }
        return {}
      })
      .reduce((prev, v) => ({ ...prev, ...v }), {})
  }

  css(css: CSSObject) {
    this.cssobject = { ...this.cssobject, ...css }
    return this
  }

  to() {
    return this.cssobject
  }
}

export const useStyles = <A extends string>(
  createStyles: (theme: (args?: Partial<ThemedStyleArgs>) => ThemedStyle) => Record<A, ThemedStyle>
): Record<A, SerializedStyles> => {
  const theme = useTheme()
  const styles = createStyles((args) => new ThemedStyle(theme, args ?? {}))

  return Object.fromEntries(
    Object.entries<ThemedStyle>(styles).map(([k, v]) => [k, css(v.to())])
  ) as Record<A, SerializedStyles>
}
