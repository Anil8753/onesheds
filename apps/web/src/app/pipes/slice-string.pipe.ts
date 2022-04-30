import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'sliceString',
})
export class SliceStringPipe implements PipeTransform {
  //
  transform(value: string[], index: number, delm: string): string {
    let final = value
      .filter((e, i) => {
        return i < index && e;
      })
      .reduce((v, e) => {
        return v + e + delm + ' ';
      }, '');

    return final.slice(0, final.length - 2);
  }
}
