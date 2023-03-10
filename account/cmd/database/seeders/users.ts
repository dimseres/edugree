import {ISeeder} from "./ISeeder";
import {PrismaClient} from "@prisma/client";

const crypto = require('crypto');
const salt = "salt";

export class UserSeeder implements ISeeder {
    seed(client: PrismaClient) {
        return async () => {
            console.log("CREATING USERS");
            await crypto.scrypt("admin", salt, 32, {N: 2048, r: 4, p: 2}, async (err: any, hash: any) => {
                if (err) {
                    throw err
                }

                const passwordSalt = crypto.scryptSync("admin", hash, 32,{N: 16384, r: 8, p: 1})
                const hashedPassword = Array.from(passwordSalt, function (byte: number) {
                    return ('0' + (byte & 0xFF).toString(16)).slice(-2);
                }).join('');
                console.log('______________', hashedPassword)
                await client.user.upsert({
                    update: {},
                    where: {email: "admin@example.com"},
                    create: {
                        email: "admin@example.com",
                        password: hashedPassword,
                        password_reset_code: null,
                        phone: "",
                        full_name: "Superadmin",
                        avatar: "https://i.pinimg.com/736x/a6/e4/41/a6e4411758a5fb005672c0e363c9def2--ewok-shih-tzu.jpg",
                        bio: "Superadmin and dot",
                        active: true,
                        created_at: new Date().toISOString(),
                        updated_at: new Date().toISOString(),
                        deleted_at: null,
                    }
                })
            });
        };
    }
}