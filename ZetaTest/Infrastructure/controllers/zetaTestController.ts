import * as express from 'express';
import { ZetaTest } from '../../Domain';
import { AddZetaTest, GetZetaTest, RemoveZetaTest, UpdateZetaTest } from '../../Application';
import { ZetaTestCriteria, ZetaTestRepository } from '../persistence/typeorm';

const router = express.Router();
const zetaTestRoute = '/zeta-test';

router.get(zetaTestRoute, async (req: express.Request, res: express.Response) => {
    try {
        const getZetaTest = new GetZetaTest(new ZetaTestRepository(), new ZetaTestCriteria(req.query as { filter: string } ));
        const zetaTest = await getZetaTest.dispatch();

        res.json(zetaTest);
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.post(zetaTestRoute, async (req: express.Request, res: express.Response) => {
    try {
        const addZetaTest = new AddZetaTest(
            new ZetaTestRepository(),
            new ZetaTest(req.body),
        );
        await addZetaTest.dispatch();

        res.status(201).json({ message: 'ZetaTest create success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.put(zetaTestRoute, async (req: express.Request, res: express.Response) => {
    try {
        const updateZetaTest = new UpdateZetaTest(
            new ZetaTestRepository(),
            new ZetaTest(req.body),
        );
        await updateZetaTest.dispatch();

        res.json({ message: 'ZetaTest update success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.delete(zetaTestRoute, async (req: express.Request, res: express.Response) => {
    try {
        const removeZetaTest = new RemoveZetaTest(new ZetaTestRepository(), req.query.id as string);
        await removeZetaTest.dispatch();

        res.json({ message: 'ZetaTest remove success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

export { router as ZetaTestRouter };
