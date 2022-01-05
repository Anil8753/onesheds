export interface UserRegistrationData {
   docType: string;
   uniqueId: string;

   firstName: string;
   lastName: string;
   email: string;
   phone: string;
   address: string;
   pincode: string;
   city: string;
   district: string;
   state: string;

   pancard: string;
   aadharcard: string;
}

export const DummyRegData = () => {
   return {
      docType: '',
      uniqueId: '',

      firstName: '',
      lastName: '',
      email: '',
      phone: '',
      address: '',
      pincode: '',
      city: '',
      district: '',
      state: '',

      pancard: '',
      aadharcard: '',
   };
};
