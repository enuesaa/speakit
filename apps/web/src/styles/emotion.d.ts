import '@emotion/react'
import { CSSObject } from '@emotion/react'

export type SurfDef = {
  main: CSSObject;
  sub: CSSObject;
  reverse: CSSObject;
  transparent: CSSObject;
}
export type SizeDef = {
  x1: CSSObject;
  x2: CSSObject;
  x3: CSSObject;
}
export type DecorateDef = {
  shadow: CSSObject;
  rounded: CSSObject;
  card: CSSObject;
}
export type HoverDef = {
  shadow: CSSObject;
}
export type AroundDef = {
  x1: CSSObject;
  x1tb: CSSObject;
  x2: CSSObject;
  x2tb: CSSObject;
  x3: CSSObject;
  x3tb: CSSObject;
}

declare module '@emotion/react' {
  export interface Theme {
    surf: SurfDef;
    size: SizeDef;
    decorate: DecorateDef;
    hover: HoverDef;
    around: AroundDef;
  }
}
