import {ISeeder} from "./ISeeder";
import {PrismaClient} from "@prisma/client";

const crypto = require('crypto');
const salt = "awnid";

export class UserSeeder implements ISeeder {
    seed(client: PrismaClient) {
        return async () => {
            console.log("CREATING USERS");
            await crypto.scrypt("admin", salt, 32, { N: 2048 },async (err: any, hash: any) => {
                if (err) {
                    throw err
                }
                const hashedPassword = hash.toString()
                await client.user.upsert({
                    update: {},
                    where: { email: "admin@example.com"},
                    create: {
                        email: "admin@example.com",
                        password: hashedPassword,
                        password_reset_code: null,
                        phone: "",
                        full_name: "Superadmin",
                        avatar: "https://i.pinimg.com/736x/a6/e4/41/a6e4411758a5fb005672c0e363c9def2--ewok-shih-tzu.jpg",
                        bio: "Superadmin and dot",
                        role_id: 1,
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