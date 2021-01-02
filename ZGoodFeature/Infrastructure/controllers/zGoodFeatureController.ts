import * as express from 'express';
import { ZGoodFeature } from '../../Domain';
import { AddZGoodFeature, GetZGoodFeature, RemoveZGoodFeature, UpdateZGoodFeature } from '../../Application';
import { ZGoodFeatureCriteria, ZGoodFeatureRepository } from '../persistence/mongoose';

const router = express.Router();
const zGoodFeatureRoute = '/z-good-feature';

router.get(zGoodFeatureRoute, async (req: express.Request, res: express.Response) => {
    try {
        const getZGoodFeature = new GetZGoodFeature(new ZGoodFeatureRepository(), new ZGoodFeatureCriteria(req.query));
        const zGoodFeature = await getZGoodFeature.dispatch();

        res.json(zGoodFeature);
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.post(zGoodFeatureRoute, async (req: express.Request, res: express.Response) => {
    try {
        const addZGoodFeature = new AddZGoodFeature(
            new ZGoodFeatureRepository(),
            new ZGoodFeature(req.body),
        );
        await addZGoodFeature.dispatch();

        res.status(201).json({ message: 'ZGoodFeature create success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.put(zGoodFeatureRoute, async (req: express.Request, res: express.Response) => {
    try {
        const updateZGoodFeature = new UpdateZGoodFeature(
            new ZGoodFeatureRepository(),
            new ZGoodFeature(req.body),
        );
        await updateZGoodFeature.dispatch();

        res.json({ message: 'ZGoodFeature update success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

router.delete(zGoodFeatureRoute, async (req: express.Request, res: express.Response) => {
    try {
        const removeZGoodFeature = new RemoveZGoodFeature(new ZGoodFeatureRepository(), req.query.id);
        await removeZGoodFeature.dispatch();

        res.json({ message: 'ZGoodFeature remove success' });
    } catch (error) {
        res.status(error.statusCode || 503).json({ error: error.message || error });
    }
});

export { router as ZGoodFeatureRouter };
