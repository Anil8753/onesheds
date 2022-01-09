import { FormGroup } from '@angular/forms';

export const getDirtyValues = (cg: FormGroup) => {
   const dirtyValues = Object();
   Object.keys(cg.controls).forEach((c) => {
      const currentControl = cg.get(c);

      if (!!currentControl && currentControl.dirty) {
         dirtyValues[c] = currentControl.value;
      }
   });

   return dirtyValues;
};
