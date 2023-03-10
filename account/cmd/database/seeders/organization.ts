// import {ISeeder} from "./ISeeder";
// import {PrismaClient} from "@prisma/client";
//
// const bcrypt = require('bcrypt');
// const bcryptRounds = 11;
//
// export class OrganizationSeeder implements ISeeder {
//     seed(client: PrismaClient) {
//         return async () => {
//             await client.organization.upsert({
//                 update: {},
//                 where: {email: "company@edugree.net"},
//                 create: {
//                     email: "organization@edugree.net",
//                     title: "Какой-то Университет",
//                     description: "lorem ipsum",
//                     avatar: "https://cdn.logo.com/hotlink-ok/logo-social.png",
//                     created_at: new Date().toISOString(),
//                     updated_at: new Date().toISOString(),
//                     deleted_at: null,
//                 }
//             })
//         };
//     }
// }