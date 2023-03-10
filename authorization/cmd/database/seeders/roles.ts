import {ISeeder} from "./ISeeder";
import {PrismaClient} from "@prisma/client";

const roles = [
    {
        name: "Creator",
        slug: "creator",
        description: "Создатель организациии умеет все",
    },
    {
        name: "Administrator",
        slug: "administrator",
        description: "Админ может создавать пользователе с ролями ниже, редактировать контент, но не может удалить организацию",
    },
    {
        name: "Moderator",
        slug: "moderator",
        description: "Редактирует ветки, дисциплины",
    },
    {
        name: "Teacher",
        slug: "teacher",
        description: "Может управлять только назначенной дисциплиной не может удалить или назначать ответственных",
    },
    {
        name: "Student",
        slug: "student",
        description: "Студент",
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