import {PrismaClient} from "@prisma/client";
import {ISeeder} from "./ISeeder";
import {UserSeeder} from "./users";
import {RolesSeeder} from "./roles";


const prisma = new PrismaClient();

const seedList: ISeeder[] = [
    new RolesSeeder,
    new UserSeeder
];

async function main() {
    for (const seeder of seedList) {
        const seedFn = seeder.seed(prisma)
        if (Array.isArray(seedFn)) {
            for (const fn of seedFn) {
                await fn();
            }
        } else {
            await seedFn();
        }
    }
}

main().then(async () => {
    await prisma.$disconnect()
}).catch(async (e) => {
    console.error(e);
    await prisma.$disconnect()
    process.exit(1);
});