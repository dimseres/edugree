import {ISeeder} from "./ISeeder";
import {PrismaClient} from "@prisma/client";

const roles = [
    {
        name: "Superadmin",
        slug: "superadmin",
        description: "Super user for create another users",
    },
    {
        name: "Moderator",
        slug: "moderator",
        description: "Super user for create another users",
    },
    {
        name: "Teacher",
        slug: "teacher",
        description: "Super user for create another users",
    },
    {
        name: "Student",
        slug: "student",
        description: "Super user for create another users",
    },
];

export class RolesSeeder implements ISeeder{
    seed(client: PrismaClient) {
        const payload = [];
        for (const roleData of roles) {
            const inner = {
                ...roleData,
                created_at: new Date().toISOString(),
                updated_at: new Date().toISOString(),
                deleted_at: null,
            }
            payload.push(async () => {
                console.log("CREATING ROLES");
                return await client.role.upsert({
                    where: { slug: inner.slug },
                    update: {},
                    create: inner
                });
            })
        }
        return payload;
    }
}