import {ISeeder} from "./ISeeder";
import {PrismaClient} from "@prisma/client";

const bcrypt = require('bcrypt');
const bcryptRounds = 11;

export class UserSeeder implements ISeeder {
    seed(client: PrismaClient) {
        return async () => {
            console.log("CREATING USERS");
            await bcrypt.hash("admin", bcryptRounds).then(async (hash: any) => {
                await client.user.upsert({
                    update: {},
                    where: { email: "admin@example.com"},
                    create: {
                        email: "admin@example.com",
                        password: hash,
                        password_reset_code: null,
                        phone: "",
                        full_name: "Superadmin",
                        avatar: "https://i.pinimg.com/736x/a6/e4/41/a6e4411758a5fb005672c0e363c9def2--ewok-shih-tzu.jpg",
                        bio: "Superadmin and dot",
                        role_id: 1,
                        active: true,
                        created_at: new Date().toISOString(),
                        updated_at: new Date().toISOString(),
                        deleted_at: new Date().toISOString(),
                    }
                })
            });
        };
    }
}